package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// WaitGroups可以通过阻塞主函数避免主协程退出
	var wg sync.WaitGroup
	//设置主函数需要等待完成的协程数为1
	wg.Add(2)

	// 协程 A
	go func() {
		fmt.Println("1")
		time.Sleep(100)
		fmt.Println("2")
		time.Sleep(100)
		fmt.Println("3")
		wg.Done()
	}()

	// 协程 B
	go func() {
		fmt.Println("A")
		time.Sleep(100)
		fmt.Println("B")
		time.Sleep(100)
		fmt.Println("C")
		wg.Done() // 前协程执行完成（Done()做的工作其实就是把需要等待的协程个数减1
	}()
	//当需要等待的协程数为0时，则不需要再等待，继续执行以下的代码
	wg.Wait()
}
