package main

import (
	"bytes"
	netMsg "distribute-chat/proto-gen-go/netmsg/pb"
	"distribute-chat/server/conf"
	dmnet "distribute-chat/server/net"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"
)

func TestProto1(t *testing.T) {

	test := &netMsg.Person{
		Name:  "a",
		Id:    10,
		Email: "realdm99@google.com",
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marsharing error")
	}

	log.Println("bytes size : ", len(data))

	msgText := "msgt"

	var tmpBuff bytes.Buffer

	binary.Write(&tmpBuff, binary.LittleEndian, 4+int32(len(data))+int32(len(msgText)))
	binary.Write(&tmpBuff, binary.LittleEndian, []byte(msgText))
	binary.Write(&tmpBuff, binary.LittleEndian, data)

	log.Println("buff size : ", tmpBuff.Len())

	bytesSize := tmpBuff.Bytes()[:4]

	var value int32

	value |= int32(bytesSize[0])
	value |= int32(bytesSize[1]) << 8
	value |= int32(bytesSize[2]) << 16
	value |= int32(bytesSize[3]) << 24

	log.Println("readSize : ", value, ", buff size : ", tmpBuff.Len())

	var size int32
	var msgByte = make([]byte, 4)
	var msgRead string

	binary.Read(&tmpBuff, binary.LittleEndian, &size)
	binary.Read(&tmpBuff, binary.LittleEndian, msgByte)

	log.Println("remain buff size : ", tmpBuff.Len())

	msgRead = string(msgByte)
	var msgBody = make([]byte, size-8)

	binary.Read(&tmpBuff, binary.LittleEndian, msgBody)

	var readProto netMsg.Person
	proto.Unmarshal(msgBody, &readProto)

	log.Println("readMsg : ", msgRead)
	log.Println("proto : ", readProto)

	log.Println("remain buff size : ", tmpBuff.Len())

}

func TestSlice(t *testing.T) {
	s := make([]int, 10)

	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5

	tt := s[:2]

	log.Println("tt = ", tt)
	log.Println("ss = ", s)
}

func TestWriteByteBuffer(t *testing.T) {
	s := dmnet.NewServer()

	s.Listen(":10055")
	s.Start()

	conn, err := net.Dial("tcp", "localhost:10055")

	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	test := &netMsg.Person{
		Name:  "a",
		Id:    10,
		Email: "realdm99@google.com",
	}

	msg := netMsg.Message_REQ_WAY

	var network bytes.Buffer

	for i := 0; i < 1; i++ {

		test.Id++

		data, err := proto.Marshal(test)

		if err != nil {
			log.Fatal("marsharing error")
		}

		binary.Write(&network, binary.LittleEndian, 4+4+int32(len(data)))
		binary.Write(&network, binary.LittleEndian, int32(msg))
		binary.Write(&network, binary.LittleEndian, data)

		//log.Println("network buffer size : ", network.Len())

		conn.Write(network.Bytes())
		network.Reset()
	}

	time.Sleep(1 * time.Second)
}

func TestJsonConfig(t *testing.T) {
	bytes, err := ioutil.ReadFile("../conf/conf_server.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	configurationSave := &conf.Configuration{
		HubInfos: []conf.HubInfo{
			{
				Name:    "hubName",
				Address: "hubAddress",
				Type:    "hubType",
			},
		},
		MsgQueue: conf.MessageQueue{
			Endpoint: "localhost",
			Port:     1000,
			User:     "user",
			Pass:     "pass",
		},
	}

	jsonBytes, _ := json.Marshal(configurationSave)
	fmt.Println(string(jsonBytes))

	configuration := conf.Configuration{}
	json.Unmarshal(bytes, &configuration)
	fmt.Println(configuration) // output: [UserA, UserB]
}

//
//func TestPacket(t *testing.T) {
//
//	test := &tutorial.Person{
//		Name :"shin dong myung",
//		Id : 100,
//		Email : "realdm99@gmail.com",
//	}
//
//	data, err := proto.Marshal(test)
//
//	if err != nil {
//		log.Fatal("marshal error")
//	}
//
//	packet := dmnet.CreatePacket("test_msg", data)
//
//	var network bytes.Buffer
//
//	enc := gob.NewEncoder(&network)
//	dec := gob.NewDecoder(&network)
//
//	err = enc.Encode(packet)
//
//	if err != nil {
//		log.Fatal("encode error : ", err)
//	}
//
//	var receivePacket dmnet.Packet
//	err = dec.Decode(&receivePacket)
//
//	if err != nil {
//		log.Fatal("decode error : ", err)
//	}
//
//	fmt.Printf("%q", receivePacket.Msg)
//}
