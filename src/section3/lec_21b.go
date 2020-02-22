package main

import "time"

func main() {
	go InfiniteLoopSleep1Sec("a")
	time.Sleep(3 * time.Second)
	go InfiniteLoopSleep1Sec("b")
}
