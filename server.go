package main

import (
	"cacher/cache"
	"fmt"
	"log"
	"net"
)

type ServerOpt struct {
	ListenAdr string
	IsLeader  bool
}

type Server struct {
	ServerOpt
	cache cache.Cacher
}

func NewServer(opts ServerOpt, c cache.Cacher) *Server {
	return &Server{
		ServerOpt: opts,
		cache:     c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAdr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}
	fmt.Println("Serever starting on ", s.ListenAdr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept  error: %s", err)
			continue
		}
		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(con net.Conn) {
	defer func() {
		con.Close()
	}()
	buf := make([]byte, 2048)
	for {
		n, err := con.Read(buf)
		if err != nil {
			log.Printf("conn read error %s", err)
			break
		}
		go s.handleCommand(con, buf[:n])
	}
}

func (s *Server) handleCommand(conn net.Conn, raw []byte) {
	msg, err := parseCommand(raw)
	if err != nil {
		fmt.Println("failed to parse command", err)
		return
	}
	switch msg.Cmd {
	case CMDset:
		if err := s.handleSetCommand(conn, msg); err != nil {
			return
		}

	}

}

func (s *Server) handleSetCommand(con net.Conn, msg *Message) error {
	fmt.Println("handlig the set command" , msg)
	return nil
}
