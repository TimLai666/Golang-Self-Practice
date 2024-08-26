package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:3306"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println(tcpAddr)

	myConn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err1)
	fmt.Println("myConn:")
	typeof(myConn)

	_, err = myConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := io.ReadAll(myConn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func typeof(v interface{}) {
	fmt.Printf("type is: %T\n", v)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
