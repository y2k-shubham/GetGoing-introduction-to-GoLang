package main

import (
	"fmt"
	"time"
)

func InfiniteLoopSleep1Sec(caller string) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("puchi-1s: %s\n", caller)
	}
}

func InfiniteLoopSleep2Sec(caller string) {
	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("puchi-2s: %s\n", caller)
	}
}

