package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("使用者退出")
			break
		}

		conent, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			panic(err)
		}
		fmt.Printf("发送了 %d 字节的数据\n", conent)
	}
}

func main() {
	Client()
}
