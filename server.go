package main

import (
	"cacher/cache"
	"context"
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
	var err error
	msg, err := parseCommand(raw)
	if err != nil {
		fmt.Println("failed to parse command", err)
		return
	}
	switch msg.Cmd {
	case CMDset:
		err = s.handleSetCommand(conn, msg)
	case CMDget:
		err = s.handleGetCommand(conn, msg)
	case CMDdel:
		err = s.handeDelCommand(conn, msg)
	case CMDhas:
		err = s.handleHasCommand(conn, msg)

	}
	if err != nil {
		fmt.Println("failed to handle command", err)
		return
	}
}

func (s *Server) handleGetCommand(con net.Conn, msg *Message) error {
	val, err := s.cache.Get(msg.Key)
	if err != nil {
		return err
	}
	_, err = con.Write(val)
	if err != nil {
		return err
	}
	go s.sendToFollowers(context.TODO(), msg)
	return nil
}

func (s *Server) handleHasCommand(con net.Conn, msg *Message) error {
	val := s.cache.Has(msg.Key)
	var err error
	if val {
		_, err = con.Write([]byte("true"))
	}
	if err != nil {
		return err
	}
	go s.sendToFollowers(context.TODO(), msg)
	return nil
}


func (s *Server) handeDelCommand(con net.Conn, msg *Message) error {
	err := s.cache.Delete(msg.Key)
	if err != nil {
		return err
	}
	go s.sendToFollowers(context.TODO(), msg)
	return nil
}

func (s *Server) handleSetCommand(con net.Conn, msg *Message) error {
	if err := s.cache.Set(msg.Key, msg.Value, msg.TTL); err != nil {
		return err
	}
	go s.sendToFollowers(context.TODO(), msg)
	return nil
}

func (s *Server) sendToFollowers(ctx context.Context, msg *Message) error {
	return nil
}
