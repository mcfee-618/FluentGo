package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := new(int)
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(a)
	fmt.Println(*a)
}
