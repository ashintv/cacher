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
		fmt.Println("Cache Server Started")
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------------------------------------")
		conn.Write([]byte("SET Foo Bar 2500000000"))
		time.Sleep(time.Second * 2)

		fmt.Println("------------------------------------")
		conn.Write([]byte("GET Foo"))
		buf := make([]byte, 1000)
		n, _ := conn.Read(buf)
		fmt.Println("Response for SET", string(buf[:n]))


		time.Sleep(time.Second * 2)
		fmt.Println("------------------------------------")
		conn.Write([]byte("HAS Foo"))
		buf = make([]byte, 1000)
		n, _ = conn.Read(buf)
		fmt.Println("Response for HAS", string(buf[:n]))


		time.Sleep(time.Second * 2)
		fmt.Println("------------------------------------")
		fmt.Println("Sending DEL")
		conn.Write([]byte("DEL Foo"))


		time.Sleep(time.Second * 2)
		fmt.Println("------------------------------------")

		conn.Write([]byte("HAS Foo"))
		buf = make([]byte, 1000)
		n, _ = conn.Read(buf)
		fmt.Println("Response for HAS", string(buf[:n]))

		
		time.Sleep(time.Second * 2)
		fmt.Println("------------------------------------")

	}()
	cacher := cache.New()
	server := NewServer(opts, cacher)
	server.Start()
}
