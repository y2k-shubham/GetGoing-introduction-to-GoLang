package main

import (
	"fmt"
	"strings"
)

func main() {

	// part1
	var m1 string
	m1 = "puchi puchi"
	fmt.Println(m1)

	// part2
	m2 := "kuchi kuchi"
	fmt.Println(m2)

	// part3
	m1 = "monchi monchi"
	fmt.Println(m1)

	// part4
	var m3 string = "kuchi"
	fmt.Println(strings.Contains(m3, m2))
	fmt.Println(strings.Contains(m2, m3))

	// part5
	fmt.Println(strings.ReplaceAll(m2, "puchi", "poncho"))
	fmt.Println(strings.ReplaceAll(m2, "kuchi", "poncho"))

	// part6
	fmt.Println(strings.Split(m1, " "))

	// part7
	fmt.Println(m1 + m2)
}
