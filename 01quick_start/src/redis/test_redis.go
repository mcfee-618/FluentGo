    package main
    import (
     "fmt"
     "github.com/gomodule/redigo/redis"
    )
    func main() {
        c, err := redis.Dial("tcp", "127.0.0.1:6379")
        if err != nil {
            fmt.Println("Connect to redis error", err)
            return
        }
        fmt.Println("redis connect succ")
        defer c.Close()
    }