package main

import "fmt"

func swap(m1 *int, m2 *int) {
	var tmp int = * m1
	* m1 = * m2
	* m2 = tmp
}

func main() {

	// part1
	var m1 int
	m1 = 2
	fmt.Println(m1)

	// part2
	var (
		m2 int = 1
		m3     = 2
	)
	fmt.Println(m2 + m3)

	// part3
	var m4 int32
	var m5 int64
	fmt.Println(int64(m4) + m5)

	// part4
	m6 := 6
	var m7 int = 7
	fmt.Println(m6 + m7)

	// part5
	var v1 int = 15
	var v2 int = 16
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
	swap(&v1, &v2)
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}
