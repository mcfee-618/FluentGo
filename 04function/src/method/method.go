package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name    string
	age     uint64
	address string
}

func (obj Person) set(key string, value string) {
	if key == "name" {
		obj.name = value
	} else if key == "address" {
		obj.address = value
	} else {
		rs, error := strconv.ParseUint(value, 10, 64)
		if error == nil {
			obj.age = rs
		}
	}
}

func main() {
	fmt.Println("start")
	p := Person{name: "fei", age: 23, address: "fei"}
	p.set("name", "shishi")
	fmt.Println(p.name)
}
