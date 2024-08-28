package main

import (
	"log"
	"net"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received %s from %s\n", buf[:n], addr)

		_, err = conn.WriteToUDP([]byte("Hello from server"), addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
