package main

import "fmt"

func swap2(m1 * int, m2 * int) {
	var tmp int
	tmp = *m1
	*m1 = *m2
	*m2 = tmp
}

func main() {
	var m1 int = 34
	var m1_ptr *int = &m1
	fmt.Printf("m1 = %d\n&m1 = %p\n*(&m1) = %d\n", m1, m1_ptr, *m1_ptr)

	m2, m3 := -19, 82
	fmt.Printf("Before swapping\nm2 = %d\nm3 = %d\n", m2, m3)
	swap2(&m2, &m3)
	fmt.Printf("After  swapping\nm2 = %d\nm3 = %d\n", m2, m3)
}
