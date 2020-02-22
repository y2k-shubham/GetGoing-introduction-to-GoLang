package main

import (
	"fmt"
	"sync"
	"time"
)

func IterativePrinter(ident string, iter int, sleepMillis int) {
	for i := 1; i <= iter; i++ {
		fmt.Printf("%s: i=%d\n", ident, i)
		time.Sleep(time.Duration(sleepMillis) * time.Millisecond)
	}
}

func UnlockedCaller(ident string, iter int, sleep int) {
	IterativePrinter(ident, iter, sleep)
}

func LockedCaller(ident string, iter int, sleep int, mutex *sync.Mutex) {
	mutex.Lock()
	IterativePrinter(ident, iter, sleep)
	mutex.Unlock()
}

func main() {
	println("UnlockedCaller")
	go UnlockedCaller("unlocked_a", 10, 2)
	go UnlockedCaller("unlocked_b", 10, 2)
	time.Sleep(time.Second * 1)

	println("LockedCaller")
	var mutex *sync.Mutex = &sync.Mutex{}
	go LockedCaller("locked_a", 10, 2, mutex)
	go LockedCaller("locked_b", 10, 2, mutex)
	time.Sleep(time.Second * 1)
}
