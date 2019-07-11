package netHandler

import (
	"log"
	"net"
)

func OnAccept(conn net.Conn) {
	log.Println("onAccept")
}