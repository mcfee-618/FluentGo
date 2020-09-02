package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err == nil {
		conn.Do("SET", "NAME", "VALUE1")
		rs, _ := conn.Do("GET", "NAME")
		value, _ := redis.String(rs, nil)
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}
}
