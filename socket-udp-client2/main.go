package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8023")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				log.Fatal(err)
			}
			conn.Write(str[:n])
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
