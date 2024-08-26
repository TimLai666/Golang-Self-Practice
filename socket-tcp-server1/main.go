package main

import (
	"fmt"
	"log"
	"net"
)

func Server() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Accepted connection:%v ip=%s\n", conn, conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		fmt.Printf("waiting %s sent data\n", conn.RemoteAddr())
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("received data:%v\n", string(buf[:n]))
	}
}

func main() {
	Server()
}
