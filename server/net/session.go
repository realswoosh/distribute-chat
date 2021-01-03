package net

import (
	"bytes"
	pb "distribute-chat/proto-gen-go/netmsg/pb"
	"distribute-chat/server/chat"
	"distribute-chat/server/conf"
	"distribute-chat/server/net/handler"
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
		recvBuff: make([]byte, conf.SessionBufferSize),
	}

	s.msgHandler = handler.MsgHandler{
		Dispatcher: map[pb.Message_Type]handler.MsgFunc{
			pb.Message_REQ_WAY: s.msgHandlerReqWay,
			pb.Message_REQ_ROOM_LIST: s.msgHandlerReqRoomList,
			pb.Message_REQ_ROOM_JOIN: s.msgHandlerReqRoomJoin,
		}}
	return s
}

func (session* Session) disconnect() {
	if session.conn != nil {
		_ = session.conn.Close()
	}
	session.manager.Remove(session)
}

func (session* Session) Send(msg pb.Message_Type, data proto.Message) {
	var marshal []byte

	if data != nil {
		convert, err := proto.Marshal(data)
		marshal = convert
		if err != nil {
			log.Fatal("Marshal error")
		}
	}

	var networkBuffer bytes.Buffer

	_ = binary.Write(&networkBuffer, binary.LittleEndian, int32(HeaderSize+int32(len(marshal))))
	_ = binary.Write(&networkBuffer, binary.LittleEndian, int32(msg))

	if data != nil {
		_ = binary.Write(&networkBuffer, binary.LittleEndian, marshal)
	}

	session.conn.Write(networkBuffer.Bytes())
}

func (session* Session) Serve() {
	defer session.disconnect()

	session.Send(pb.Message_WAY, nil)

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

		if session.recvSize > conf.SessionBufferSize {
			log.Printf("disconnect session. recv buffer overflow disconnect %v", session.conn.RemoteAddr())
			session.disconnect()
			return
		}

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

				val, exists := session.msgHandler.Dispatcher[pb.Message_Type(msg)]

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
	var msgProto pb.Person

	proto.Unmarshal(msgBody, &msgProto)

	log.Println("proto : " , msgProto)
}

func (session* Session)msgHandlerReqWay(msgBody []byte) {

	var msgReqWay pb.ReqWay

	proto.Unmarshal(msgBody, &msgReqWay)

	session.Info.uid, _ = uuid.NewUUID()
	session.Info.Name = msgReqWay.Name
	session.Info.Email = msgReqWay.Email

	msgAckWay := pb.AckWay{
		Name : session.Info.Name,
		Email : session.Info.Email,
		Uuid : session.Info.uid.String(),
	}

	session.Send(pb.Message_ACK_WAY, &msgAckWay)
}

func (session* Session)msgHandlerReqRoomList(msgBody []byte) {
	rooms := pb.AckRoomList{}
	for _, r := range chat.RoomManagerInstance().List() {
		msgRoomInfo := &pb.RoomInfo{
			RoomIdx:              r.Idx,
			Title:                r.Title,
			ParticipateCount:     r.ParticipateCount,
		}
		rooms.RoomList = append(rooms.RoomList, msgRoomInfo)
	}

	session.Send(pb.Message_ACK_ROOM_LIST, &rooms)
}

func (session* Session)msgHandlerReqRoomJoin(msgBody []byte) {
	var msgReqRoomJoin pb.ReqRoomJoin

	_ = proto.Unmarshal(msgBody, &msgReqRoomJoin)
}