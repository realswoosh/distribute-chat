package net

import (
	"io"
	"log"
	"net"
)

type Session struct {
	conn net.Conn
	manager *SessionManager
	recvBuff []byte
}

func CreateSession(manager *SessionManager, conn net.Conn) *Session {
	return &Session{
		conn: conn,
		manager: manager,
		recvBuff: make([]byte, 8192),
	}
}

func (session* Session) disconnect() {
	session.manager.Remove(session)
}

func (session* Session) Serve() {
	defer session.disconnect()

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
		if n > 0 {
			data := session.recvBuff[:n]
			log.Printf(string(data))
		}
	}
}