package main

import (
	"../proto-gen-go"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestProto1(t *testing.T) {

	test := &tutorial.Person{
		Name : "a",
		Id : 10,
		Email : "realdm99@google.com",
	}

	data, err := proto.Marshal(test)
	if (err != nil) {
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

func TestNetworkReceived(t *testing.T) {
	//s := dmnet.NewServer()
	//
	//s.Listen(":10055")
	//s.Start()
	//
	//conn, err := net.Dial("tcp", "localhost:10055")
	//
	//if nil != err {
	//	log.Fatalf("failed to connect to server")
	//}

	test := &tutorial.Person{
		Name : "a",
		Id : 10,
		Email : "realdm99@google.com",
	}


	data, err := proto.Marshal(test)

	if err != nil {
		log.Fatal("marsharing error")
	}

	size := len(data)

	log.Println("data size = ", size);

	var sizeByte []byte;
	sizeByte := make([]byte, 4)
	binary.LittleEndian.
	

}