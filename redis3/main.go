package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Send("SET", "foo", "bar")
	c.Send("SET", "foo2", "bar2")
	c.Flush()

	v, err := c.Receive()
	fmt.Printf("v:%v, err:%v\n", v, err)

	v, err = c.Receive()
	fmt.Printf("v:%v, err:%v\n", v, err)

	v, err = c.Receive()
	fmt.Printf("v:%v, err:%v\n", v, err)
}
