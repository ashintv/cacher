package main

import (
	"cacher/cache"
	"fmt"
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
		conn.Write([]byte("SET Foo Bar 2500000000"))
		time.Sleep(time.Second * 2)
		conn.Write([]byte("GET Foo"))

		buf := make([]byte, 1000)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))

	}()
	cacher := cache.New()
	server := NewServer(opts, cacher)
	server.Start()
}
