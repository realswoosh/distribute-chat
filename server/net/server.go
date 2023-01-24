package net

import (
	"log"
	"net"
)

type Server struct {
	listener       net.Listener
	sessionManager *SessionManager
}

func NewServer() *Server {
	s := &Server{
		sessionManager: CreateSessionManager(),
	}
	return s
}

func (s *Server) Listen(address string) error {
	ln, err := net.Listen("tcp", address)
	if err == nil {
		s.listener = ln
	} else {
		log.Println(err)
	}
	log.Printf("listening on %v", address)
	return err
}

func (s *Server) startServer() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Println(err)
		} else {
			session := s.accept(conn)
			go session.Serve()
		}
	}
}

func (s *Server) Start() {
	go s.startServer()
}

func (s *Server) Stop() {
	_ = s.listener.Close()
}

func (s *Server) accept(conn net.Conn) *Session {
	log.Printf("accepting new connection from %v, total session: %v",
		conn.RemoteAddr().String(),
		s.sessionManager.TotalSessionCount())

	session := CreateSession(s, conn)

	s.sessionManager.Append(session)

	return session
}

func (s *Server) disconnect(session *Session) {
	log.Println("disconnect session remove")
	s.sessionManager.Remove(session)
}
