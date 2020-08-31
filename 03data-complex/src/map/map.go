package main

import "fmt"

func main() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["a"] = "b"
	map1["b"] = "c"
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	_, ok := map1["a"]
	if ok {
		fmt.Println("存在d")
	} else {
		fmt.Println("不存在d")
	}
}
