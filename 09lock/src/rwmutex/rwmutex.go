package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex *sync.RWMutex = new(sync.RWMutex)

	var group *sync.WaitGroup = new(sync.WaitGroup)

	group.Add(100)

	go write(mutex, group)

	for i := 1; i < 100; i++ {
		go read(mutex, group)
	}

	group.Wait()

}

func read(mutex *sync.RWMutex, group *sync.WaitGroup) {
	mutex.RLock()
	fmt.Println("读操作")
	time.Sleep(time.Duration(3) * time.Second)
	mutex.RUnlock()
	group.Done()
}

func write(mutex *sync.RWMutex, group *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("写操作")
	time.Sleep(time.Duration(3) * time.Second)
	mutex.Unlock()
	group.Done()
}
