package main

import "fmt"

func test(ch chan int) {
	ch <- 1
	ch <- 2
}

func main() {
	ch := make(chan int)
	go test(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
