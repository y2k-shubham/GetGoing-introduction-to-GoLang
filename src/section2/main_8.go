package main

import (
	"fmt"
)

func if_else_demo(flag *bool) {
	if flag == nil {
		fmt.Println("nil")
	} else if *flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func for_loop_demo_1() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("i=%d\n", i)
	}
}

func for_loop_demo_2() {
	arr := []string{"puchi", "kuchi", "monchi"}
	for idx, value := range arr {
		fmt.Printf("arr[%d] = %s\n", idx, value)
	}
}

func for_loop_demo_3() {
	var myMap map[string]interface{} = map[string]interface{}{"k1": "v1", "k2": "v2"}
	for key, value := range myMap {
		fmt.Printf("myMap[%s] = %s\n", key, value)
	}
}

func main() {
	var flag1 bool = false
	var flag2 bool = true
	if_else_demo(&flag1)
	if_else_demo(&flag2)
	if_else_demo(nil)

	for_loop_demo_1()

	for_loop_demo_2()
	for_loop_demo_3()
}
