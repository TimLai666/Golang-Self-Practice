package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	conn.Send("HSET", "students", "name", "John", "age", "25")
	conn.Send("HSET", "students", "score", "90")
	conn.Send("HGET", "students", "age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("res2:%v \n", res2)
	res3, err := conn.Receive()
	fmt.Printf("res3:%s", res3)
}
