package main

import "fmt"

type msg struct {
	content string
}

func main() {
	fmt.Println("start")
	c := test(63, 0)
	fmt.Println(c.content)
	fmt.Println("end")
}

func test(a int32, b int32) (obj *msg) {

	obj = new(msg)
	obj.content = "111"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获异常")
			obj.content = "222"
		}
	}()
	_ = a / b
	return obj
}

//如果想在recover继续返回，需要显式的设置返回值变量的名字
