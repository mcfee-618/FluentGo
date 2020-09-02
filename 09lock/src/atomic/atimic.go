package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	num := int32(0)

	for i := int32(0); i <= 1000; {
		go func() {
			atomic.AddInt32(&i, 1)
			atomic.AddInt32(&num, 1)
			fmt.Printf("第 %d 次循环, %d \n", i, num)
			time.Sleep(30 * time.Millisecond)
		}()
	}

}
