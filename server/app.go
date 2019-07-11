package main

import (
	"log"
	"net"
	"./netHandler"
)

func main() {
	ln, err := net.Listen("tcp", "10055");
	if err == nil {
		log.Println("listen failed")
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept();
		if err == nil {
			// log
			continue;
		}

		defer conn.Close()

		go netHandler.OnAccept(conn)
	}

	log.Println("main")
}