package main

import "fmt"

func todo() {
	// part1
	var arr1 []int
	arr1 = []int{1, 2, 3, 4}
	fmt.Println(arr1)

	// part2
	arr2 := []string{"kuchi", "puchi"}
	fmt.Println(arr1, arr2)

	// part3
	arr2 = append(arr2, "poncho")
	fmt.Println(append(arr2, "monchi", "goojy"))
}

func main() {
	todo()
}
