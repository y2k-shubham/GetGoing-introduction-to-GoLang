package main

import "fmt"

func main() {
	// go run src/section2/main_3a.go src/section2/main_3b.go
	// https://stackoverflow.com/questions/14155122/how-to-call-function-from-another-file-in-go-language#comment19630642_14155156
	fmt.Println(hcf(36, 24))
	fmt.Println(hcf(36, 19))
}
