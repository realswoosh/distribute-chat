package main

import (
	"../proto-gen-go"
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