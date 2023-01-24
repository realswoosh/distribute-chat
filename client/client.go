package main

import (
	"bufio"
	"bytes"
	"distribute-chat/proto-gen-go/netmsg/pb"
	distc "distribute-chat/server/net"
	"distribute-chat/server/net/handler"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"strings"
)

var msgHandler handler.MsgHandler
var client net.Conn

func initializeHandler() {
	msgHandler = handler.MsgHandler{
		Dispatcher: map[pb.Message_Type]handler.MsgFunc{
			pb.Message_PONG:    msgHandlerPong,
			pb.Message_WAY:     msgHandlerWay,
			pb.Message_ACK_WAY: msgHandlerAckWay,
		}}
}

func main() {
	initializeHandler()

	conn, err := net.Dial("tcp", "localhost:20055")

	if nil != err {
		log.Fatal("failed to connect server")
	}

	defer conn.Close()

	client = conn

	go func(c net.Conn) {
		recvData := make([]byte, 8192)
		recvSize := int32(0)
		data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성
		for {
			n, err := c.Read(data) // 서버에서 받은 데이터를 읽음
			if err != nil {
				fmt.Println("disconnect = ", err)
				return
			}

			recvData = append(recvData[:recvSize], data[:n]...)
			recvSize += int32(n)
			readSize := int32(0)

			for recvSize-readSize >= distc.HeaderSize {
				size := distc.ByteToInt32(recvData[readSize : readSize+4])
				msgType := distc.ByteToInt32(recvData[readSize+4 : readSize+8])
				msgSize := size - distc.HeaderSize
				msgBody := recvData[readSize+distc.HeaderSize : readSize+distc.HeaderSize+msgSize]
				println("msg = ", pb.Message_Type(msgType), ", msgBody = ", string(msgBody))
				msgFunc := msgHandler.Dispatcher[pb.Message_Type(msgType)]
				msgFunc(msgBody)
				readSize += size
			}
			recvSize -= readSize
		}
	}(conn)

	isRun := true
	for isRun {
		fmt.Print("-> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("exit", text) == 0 {
			fmt.Println("bye bye!!")
			isRun = false
		}

		ping := pb.Ping{
			Name: text,
		}
		Send(conn, pb.Message_PING, &ping)
	}
}

func Send(conn net.Conn, msg pb.Message_Type, data proto.Message) {
	var marshal []byte

	if data != nil {
		convert, err := proto.Marshal(data)
		marshal = convert
		if err != nil {
			log.Fatal("Marshal error")
		}
	}

	var networkBuffer bytes.Buffer

	_ = binary.Write(&networkBuffer, binary.LittleEndian, int32(distc.HeaderSize+int32(len(marshal))))
	_ = binary.Write(&networkBuffer, binary.LittleEndian, int32(msg))

	if data != nil {
		_ = binary.Write(&networkBuffer, binary.LittleEndian, marshal)
	}

	_, _ = conn.Write(networkBuffer.Bytes())
}

func msgHandlerPong(msgBody []byte) {
	var msgPong pb.Pong
	_ = proto.Unmarshal(msgBody, &msgPong)
	println("pong = " + msgPong.Name)
}

func msgHandlerWay(msgBody []byte) {
	var msgReqWay pb.ReqWay
	msgReqWay.Name = "realdm99"
	msgReqWay.Email = "realdm99@gmail.com"
	Send(client, pb.Message_REQ_WAY, &msgReqWay)
}

func msgHandlerAckWay(msgBody []byte) {
	var msgAckWay pb.AckWay
	_ = proto.Unmarshal(msgBody, &msgAckWay)
	fmt.Println("msgAckWay = ", msgAckWay.Uuid)
}
