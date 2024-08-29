package main

import (
	"fmt"
	"net"

	"github.com/HazelnutParadise/Go-Utils/errutil"
	"github.com/HazelnutParadise/Go-Utils/timeutil"
)

type Heartbeat struct {
	endTime int64
}

var connMap map[net.Conn]*Heartbeat

func handleConn(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			break
		}

		if timeutil.UnixAfterSeconds(0) > connMap[conn].endTime {
			fmt.Println("Connection timed out")
			break
		} else {
			connMap[conn].endTime = timeutil.UnixAfterSeconds(5)
		}

		if string(buffer[:n]) == "1" {
			conn.Write([]byte("1"))
			continue
		}

		for con, heart := range connMap {
			if con == conn {
				continue
			}
			if timeutil.UnixAfterSeconds(0) > heart.endTime {
				delete(connMap, con)
				con.Close()
				fmt.Printf("Connection %s timed out\n", con.RemoteAddr())
				fmt.Printf("Remaining connections: %v\n", connMap)
				continue
			}
			con.Write(buffer[:n])
		}
	}
}

func main() {
	l := errutil.PanicOnErr(net.Listen, "tcp", "127.0.0.1:8086")[0].(net.Listener)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		connMap[conn] = &Heartbeat{
			endTime: timeutil.UnixAfterSeconds(5),
		}
		go handleConn(conn)
	}
}
