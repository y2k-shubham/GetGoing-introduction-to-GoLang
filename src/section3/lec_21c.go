package main

import (
	"fmt"
)

func main() {
	fmt.Println("main-c")
	go InfiniteLoopSleep1Sec("c")
	fmt.Println("main-c")
	select {}
}
