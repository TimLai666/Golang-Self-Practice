package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8012")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, raddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read from %s: %s\n", raddr, buf[:n])

	_, err = conn.WriteToUDP([]byte("hello"), raddr)
	if err != nil {
		panic(err)
	}
}
