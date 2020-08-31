package main

import "fmt"

func main() {
	var array []int32
	array = make([]int32, 10)
	array = append(array, 122)
	array[0] = -222
	fmt.Println(array[0])
	fmt.Println(array[10]) // 使用原有的容量来分配一个新元素
	fmt.Println(array[0:2])
	fmt.Println(len(array))
	for index, num := range array {
		fmt.Println(index, num)
	}
}
