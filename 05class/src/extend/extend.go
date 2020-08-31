package main

import "fmt"

type A struct {
	age int32
}

type B struct {
	A
	name string
}

func main() {
	b := new(B)
	b.age = 23
	fmt.Println(b.age)
}
