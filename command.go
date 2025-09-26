package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

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


func parseCommand(raw []byte) (*Message, error) {
	var (
		rawStr = string(raw)
		parts  = strings.Split(rawStr, " ")
	)
	len_ := len(parts)
	if len_ < 2 {
		return nil, errors.New("Invalid protocol format")
	}

	message := &Message{
		Cmd: Command(parts[0]),
		Key: []byte(parts[1]),
	}
	if message.Cmd == CMDset {
		if len_ < 4 {
			return nil, errors.New("Invalid SET protocol format")
		}
		message.Value = []byte(parts[2])
		seconds, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println("invalid command Times")
		}
		message.TTL = time.Second * time.Duration(seconds)
	}
	return message, nil
}
