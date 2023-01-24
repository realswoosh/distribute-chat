package net

import (
	"bytes"
	pb "distribute-chat/proto-gen-go/netmsg/pb"
	"distribute-chat/server/chat"
	"distribute-chat/server/conf"
	"distribute-chat/server/net/handler"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
)

type SessionInfo struct {
	uid   uuid.UUID
	Name  string
	Email string
}
type Session struct {
	conn       net.Conn
	server     *Server
	recvData   []byte
	recvSize   int32
	recvBuff   []byte
	Info       SessionInfo
	msgHandler handler.MsgHandler
}

func CreateSession(server *Server, conn net.Conn) *Session {
	s := &Session{
		conn:     conn,
		server:   server,
		recvSize: 0,
		recvBuff: make([]byte, conf.SessionBufferSize),
	}

	s.msgHandler = handler.MsgHandler{
		Dispatcher: map[pb.Message_Type]handler.MsgFunc{
			pb.Message_PING:          s.msgHandlerPing,
			pb.Message_REQ_WAY:       s.msgHandlerReqWay,
			pb.Message_REQ_ROOM_LIST: s.msgHandlerReqRoomList,
			pb.Message_REQ_ROOM_JOIN: s.msgHandlerReqRoomJoin,
		}}
	return s
}

func (session *Session) disconnect() {
	if session.conn != nil {
		_ = session.conn.Close()
	}
	session.server.disconnect(session)
	session.server = nil
}

func (session *Session) Send(msg pb.Message_Type, data proto.Message) {
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

	_, _ = session.conn.Write(networkBuffer.Bytes())

	fmt.Println("send msgType = ", msg.Type())
}

func (session *Session) Serve() {
	defer session.disconnect()

	session.Send(pb.Message_WAY, nil)

	data := make([]byte, 4096)
	for {
		recvSize, err := session.conn.Read(data)
		if err != nil {
			if io.EOF == err {
				log.Printf("disconnect session. %v", session.conn.RemoteAddr())
				return
			}
			log.Printf("fail to receive data; err : %v", err)
			return
		}

		recvSize32 := int32(recvSize)

		if session.recvSize > conf.SessionBufferSize {
			log.Printf("disconnect session. recv buffer overflow disconnect %v", session.conn.RemoteAddr())
			session.disconnect()
			return
		}

		log.Println("total receive n = ", recvSize)

		session.recvData = append(session.recvData[:session.recvSize], data[:recvSize]...)

		session.recvSize += recvSize32

		readSize := int32(0)

		for recvSize32-readSize >= HeaderSize {
			sizeSlice := session.recvData[readSize : readSize+4]
			size := ByteToInt32(sizeSlice)

			log.Println("msg size = ", size)

			msgType := ByteToInt32(session.recvData[readSize+4 : readSize+8])
			msgSize := size - HeaderSize
			msgBody := session.recvData[readSize+HeaderSize : readSize+HeaderSize+msgSize]

			msgHandler, exists := session.msgHandler.Dispatcher[pb.Message_Type(msgType)]

			if !exists {
				log.Println("message not exists")
			} else {
				msgHandler(msgBody)
			}

			readSize += size
		}

		session.recvData = session.recvData[readSize:session.recvSize]
		session.recvSize = session.recvSize - readSize

		println("aaaaa")
	}
}

func (session *Session) msgHandlerTest(msgBody []byte) {
	var msgProto pb.Person

	_ = proto.Unmarshal(msgBody, &msgProto)

	log.Println("proto : ", msgProto)
}

func (session *Session) msgHandlerReqWay(msgBody []byte) {

	var msgReqWay pb.ReqWay

	_ = proto.Unmarshal(msgBody, &msgReqWay)

	session.Info.uid, _ = uuid.NewUUID()
	session.Info.Name = msgReqWay.Name
	session.Info.Email = msgReqWay.Email

	msgAckWay := pb.AckWay{
		Name:  session.Info.Name,
		Email: session.Info.Email,
		Uuid:  session.Info.uid.String(),
	}

	session.Send(pb.Message_ACK_WAY, &msgAckWay)
}

func (session *Session) msgHandlerPing(msgBody []byte) {
	var msgPing pb.Ping
	_ = proto.Unmarshal(msgBody, &msgPing)
	println("ping = " + msgPing.Name)
	var msgPong pb.Pong
	msgPong.Name = msgPing.Name
	session.Send(pb.Message_PONG, &msgPong)
}

func (session *Session) msgHandlerReqRoomList(msgBody []byte) {
	rooms := pb.AckRoomList{}
	for _, r := range chat.RoomManagerInstance().List() {
		msgRoomInfo := &pb.RoomInfo{
			RoomIdx:          r.Idx,
			Title:            r.Title,
			ParticipateCount: r.ParticipateCount,
		}
		rooms.RoomList = append(rooms.RoomList, msgRoomInfo)
	}

	session.Send(pb.Message_ACK_ROOM_LIST, &rooms)
}

func (session *Session) msgHandlerReqRoomJoin(msgBody []byte) {
	var msgReqRoomJoin pb.ReqRoomJoin

	_ = proto.Unmarshal(msgBody, &msgReqRoomJoin)
}
