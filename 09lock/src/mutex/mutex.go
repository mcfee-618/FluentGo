package main

import (
	"fmt"
	"sync"
)

var num int = 0

var mutex sync.Mutex

var wg sync.WaitGroup

func incr(wg *sync.WaitGroup) {
	mutex.Lock()
	num += 1
	mutex.Unlock()
	wg.Done()
}

func main() {

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		incr(&wg)
	}
	fmt.Println(num)
	wg.Wait()
}

// A WaitGroup must not be copied after first use
// 1.全局变量 2.使用指针传递参数
