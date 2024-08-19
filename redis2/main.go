package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("SET", "key", "value")
	if err != nil {
		panic(err)
	}
	r, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
