package net

import (
	"../../proto-gen-go"
	"../chat"
	"./handler"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
)

type SessionInfo struct {
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
		}}
	return s
}

func (session* Session) disconnect() {
	session.manager.Remove(session)
}

func (session* Session) Send(msgText string, data []byte) {
	var networkBuffer bytes.Buffer

	binary.Write(&networkBuffer, binary.LittleEndian, 4+int32(len(data))+int32(len(msgText)))
	binary.Write(&networkBuffer, binary.LittleEndian, []byte(msgText))
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
	var msgAckWay netmsg.AckWay

	proto.Unmarshal(msgBody, &msgAckWay)

	log.Println("proto : ", msgAckWay)
	session.Info.Name = msgAckWay.Name
	session.Info.Email = msgAckWay.Email
}

func (session* Session)msgHandlerReqRoomList(msgBody []byte) {
	roomInfos := chat.RoomManangerInstance().List()
	msgRoomList := netmsg.AckRoomList{}
	for _, r := range roomInfos {
		msgRoomInfo := &netmsg.RoomInfo{
			RoomIdx:              r.Idx,
			Title:                r.Title,
			ParticipateCount:     r.ParticipateCount,
		}
		msgRoomList.RoomList = append(msgRoomList.RoomList, msgRoomInfo)
	}
}
