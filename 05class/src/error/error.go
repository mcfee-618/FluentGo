package main

import (
	"errors"
	"fmt"
)

func devideNum(a int32, b int32) (rs int32, exception error) {
	if b == 0 {
		return 0, errors.New("info1")
	}
	return a % b, nil
}

func main() {
	defer fmt.Println("finally功能函数")
	_, err := devideNum(2, 0)
	if err != nil {
		fmt.Println("出错了")
	}
}
