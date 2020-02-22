package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int = make(chan int)

	go func(sendVal int) {
		fmt.Println("1: about to send")
		c <- sendVal
		fmt.Println("1: sent")
	}(14)

	// this one won't be able to send value; it will get stuck
	// but since it is invoked as a goroutine, we are good to go
	go func(sendVal int) {
		fmt.Println("2: about to send")
		c <- sendVal
		fmt.Println("2: sent")
	}(18)

	go func() {
		//time.Sleep(time.Millisecond * 10)
		var sniffVal int = <-c
		fmt.Printf("inside: sniffed value %d\n", sniffVal)
	}()

	fmt.Printf("outside: sniffed value %d\n" , <-c)

	time.Sleep(time.Second)
	println("puchi")
}
