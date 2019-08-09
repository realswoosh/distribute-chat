package net

import (
	"../../proto-gen-go"
	"../chat"
	"./handler"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
)

type SessionInfo struct {
	uid uuid.UUID
	Name string
	Email string
}
type Session struct {
	conn net.Conn
	manager *SessionManager
	recvData []byte
	recvSize int32
	recvBuff []byte
	Info SessionInfo
	msgHandler handler.MsgHandler
}

func CreateSession(manager *SessionManager, conn net.Conn) *Session {
	s := &Session{
		conn:     conn,
		manager:  manager,
		recvSize: 0,
		recvBuff: make([]byte, 4096),
	}

	s.msgHandler = handler.MsgHandler{
		Dispatcher: map[netmsg.Message_Type]handler.MsgFunc{
			netmsg.Message_REQ_WAY: s.msgHandlerReqWay,
			netmsg.Message_REQ_ROOM_LIST: s.msgHandlerReqRoomList,
		}}
	return s
}

func (session* Session) disconnect() {
	session.manager.Remove(session)
}

func (session* Session) Send(msg netmsg.Message_Type, data proto.Message) {
	marshal, err := proto.Marshal(data)
	if err != nil {
		log.Fatal("marsharing error")
	}

	var networkBuffer bytes.Buffer

	binary.Write(&networkBuffer, binary.LittleEndian, int32(HeaderSize + int32(len(marshal))))
	binary.Write(&networkBuffer, binary.LittleEndian, int32(msg))
	binary.Write(&networkBuffer, binary.LittleEndian, data)

	session.conn.Write(networkBuffer.Bytes())
}

func (session* Session) Serve() {
	defer session.disconnect()

	session.conn.Write([]byte("way"))

	for {
		n, err := session.conn.Read(session.recvBuff)
		if err != nil {
			if io.EOF == err {
				log.Printf("disconnect session. %v", session.conn.RemoteAddr())
				return
			}
			log.Printf("fail to receive data; err : %v", err)
			return
		}
		session.recvSize += int32(n)

		log.Println("receive n = ", n)

		session.recvData = append(session.recvData, session.recvBuff[:n]...)

		if int32(n) >= HeaderSize {

			sizeSlice := session.recvData[:4]

			size := byteToInt32(sizeSlice)

			log.Println("receive size = ", n, ", packet size : ", size)

			for session.recvSize >= size {

				msgTemp := session.recvData[4 : 8]
				msg := byteToInt32(msgTemp)
				msgBody := session.recvData[HeaderSize : size]

				val, exists := session.msgHandler.Dispatcher[netmsg.Message_Type(msg)]

				if !exists {
					log.Println("message not exists")
				} else {
					val(msgBody)
				}

				session.recvData = session.recvData[size:]
				session.recvSize = session.recvSize - size
			}
		}
	}
}

func (session* Session)msgHandlerTest(msgBody []byte) {
	var msgProto netmsg.Person

	proto.Unmarshal(msgBody, &msgProto)

	log.Println("proto : " , msgProto)
}

func (session* Session)msgHandlerReqWay(msgBody []byte) {

	var msgReqWay netmsg.ReqWay

	proto.Unmarshal(msgBody, &msgReqWay)

	session.Info.uid, _ = uuid.NewUUID()
	session.Info.Name = msgReqWay.Name
	session.Info.Email = msgReqWay.Email

	var msgAckWay netmsg.AckWay

	msgAckWay.Name = session.Info.Name
	msgAckWay.Email = session.Info.Email
	msgAckWay.Uuid = session.Info.uid.String()

	session.Send(netmsg.Message_ACK_WAY, &msgAckWay)
}

func (session* Session)msgHandlerReqRoomList(msgBody []byte) {
	rooms := netmsg.AckRoomList{}
	for _, r := range chat.RoomManangerInstance().List() {
		msgRoomInfo := &netmsg.RoomInfo{
			RoomIdx:              r.Idx,
			Title:                r.Title,
			ParticipateCount:     r.ParticipateCount,
		}
		rooms.RoomList = append(rooms.RoomList, msgRoomInfo)
	}

	session.Send(netmsg.Message_ACK_ROOM_LIST, &rooms)
}
