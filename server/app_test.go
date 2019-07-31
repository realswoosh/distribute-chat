package main

import (
	"../proto-gen-go"
	dmnet "./net"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"testing"
	"time"
)

func TestProto1(t *testing.T) {

	test := &tutorial.Person{
		Name : "a",
		Id : 10,
		Email : "realdm99@google.com",
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marsharing error")
	}

	fmt.Print(data)
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

func WriteByteBuffer() <- chan string{
	c := make(chan string)

	go func() {
		s := dmnet.NewServer()

		s.Listen(":10055")
		s.Start()

		conn, err := net.Dial("tcp", "localhost:10055")

		if nil != err {
			log.Fatalf("failed to connect to server")
		}

		test := &tutorial.Person{
			Name:  "a",
			Id:    10,
			Email: "realdm99@google.com",
		}

		data, err := proto.Marshal(test)

		if err != nil {
			log.Fatal("marsharing error")
		}

		packet := dmnet.CreatePacket("test_msg", data);
		var network bytes.Buffer

		binary.Write(&network, binary.LittleEndian, packet)

		//enc := gob.NewEncoder(&network)
		//enc.Encode(packet)

		log.Println("network buffer size : ", network.Len())

		conn.Write(network.Bytes())

		time.Sleep(10 * time.Second)

		c <- "done1"
	}()
	return c
}

func TestWriteByteBuffer(t *testing.T) {
	f := WriteByteBuffer()
	if <-f != "done" {

	}
}

func TestPacket(t *testing.T) {

	test := &tutorial.Person{
		Name :"shin dong myung",
		Id : 100,
		Email : "realdm99@gmail.com",
	}

	data, err := proto.Marshal(test)

	if err != nil {
		log.Fatal("marshal error")
	}

	packet := dmnet.CreatePacket("test_msg", data)

	var network bytes.Buffer

	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err = enc.Encode(packet)

	if err != nil {
		log.Fatal("encode error : ", err)
	}

	var receivePacket dmnet.Packet
	err = dec.Decode(&receivePacket)

	if err != nil {
		log.Fatal("decode error : ", err)
	}

	fmt.Printf("%q", receivePacket.Msg)
}