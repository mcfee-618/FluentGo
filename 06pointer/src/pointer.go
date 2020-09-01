package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int32
}

func main() {
	fmt.Println("start")
	p := Person{name: "fei", age: 22}
	p.age = 23
	fmt.Println(p.age)
	p1 := new(Person)
	p1.age = 88
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(p1.age)
}
