package main

import (
	"../proto-gen-go"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestProtoBuff(t *testing.T) {
	test := &tutorial.Person{
		Name: "name",
		Id: 100,
		Email: "email",
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Print("data : ", data)
}
