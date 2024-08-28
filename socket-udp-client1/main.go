package main

import (
	"net"

	"github.com/HazelnutParadise/Go-Utils/errutil"
)

func main() {
	conn := errutil.PanicOnErr(net.Dial, "udp", "127.0.0.1:8012")[0].(net.Conn)
	defer conn.Close()
	conn.Write([]byte("Hello, World!"))

	buf := make([]byte, 1024)
	n := errutil.PanicOnErr(conn.Read, buf)[0].(int)
	println(string(buf[:n]))
}
