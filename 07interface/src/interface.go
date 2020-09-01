package main

import (
	"fmt"
	"reflect"
)

type person interface {
	eat()
	sleep()
}

type Boy struct {
}

func (Boy) eat() {
	fmt.Println("eat")
}

func (Boy) sleep() {
	fmt.Println("sleep")
}

func test(p1 person) {
	fmt.Println(reflect.TypeOf(p1))
	p1.eat()
	p1.sleep()
}

func main() {
	fmt.Println("start")
	var item interface{}
	item = 1
	item = "333"
	fmt.Println(item)
	test(Boy{})
}
