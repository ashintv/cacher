package main

import "time"

type Command string

const (
	CMDset Command = "SET"
	CMDget Command = "GET"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

func ParseCommand(raw []byte) (*Message, error) {

	return nil, nil
}
