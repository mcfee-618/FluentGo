package main

import "fmt"

type Person struct {
	name string
	age  int32
}

func (p Person) show() {
	fmt.Println(p.name, p.age)
}

func main() {
	a, b := compute(1, 3)
	fmt.Println(a, b)
	person := new(Person)
	person.name = "fei"
	person1 := &Person{name: "fei", age: 2}
	fmt.Println(person1.age)
}

func compute(a int32, b int32) (int32, int32) {
	return a + b, a - b
}
