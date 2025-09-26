package main

import (
	"cacher/cache"
	"log"
	"net"
	"time"
)

func main() {
	opts := ServerOpt{
		ListenAdr: ":3000",
		IsLeader:  true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("Set Foo Bar 2500"))
	}()
	cacher := cache.New()
	server := NewServer(opts, cacher)
	server.Start()
}


