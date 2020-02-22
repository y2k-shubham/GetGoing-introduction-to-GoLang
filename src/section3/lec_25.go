package main

import (
	"fmt"
	"time"
)

func BufferedChannelDemo1() {
	var c chan int = make(chan int, 2)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("pushing i=%d-->\n", i)
			c <- i
		}
		close(c)
	}()

	time.Sleep(time.Millisecond * 1000)
	for i := range c {
		fmt.Printf("<--pulling i=%d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	BufferedChannelDemo1()
}
