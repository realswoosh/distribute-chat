package net

import (
	"io"
	"log"
	"net"
)

type Session struct {
	conn net.Conn
	manager *SessionManager
	recvData []byte
	recvSize int32
	recvBuff []byte
}

func CreateSession(manager *SessionManager, conn net.Conn) *Session {
	return &Session{
		conn: conn,
		manager: manager,
		recvSize: 0,
		recvBuff: make([]byte, 4096),
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
		session.recvSize += int32(n);

		log.Println("receive n = ", n)

		session.recvData = append(session.recvData, session.recvBuff[:n]...)

		if int32(n) >= HeaderSize {

			sizeSlice := session.recvData[:4]

			var size int32

			size |= int32(sizeSlice[0])
			size |= int32(sizeSlice[1]) << 8
			size |= int32(sizeSlice[2]) << 16
			size |= int32(sizeSlice[3]) << 24

			log.Println("receive size = ", n, ", packet size : ", size)

			if session.recvSize >= size {

				msg := string(session.recvData[4 : 8])

				log.Println("receive msg = ", msg)

				body := make([]byte, size - 8)
				copy(body, session.recvBuff[8 : size])
				CreatePacket(msg, body)
				session.recvData = session.recvData[size:]
			}
		}
	}
}