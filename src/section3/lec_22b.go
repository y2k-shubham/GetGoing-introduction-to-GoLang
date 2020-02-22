package main

import (
	"fmt"
	"sync"
)

func Printer(ident string) {
	fmt.Println("my-puchi: " + ident)
}

func NormalCaller(ident string) {
	Printer(ident)
}

func WaitReleaseCaller(ident string, group *sync.WaitGroup) {
	Printer(ident)
	group.Done()
}

func main() {
	// cleaner way for blocking than sleep / select
	var group * sync.WaitGroup = &sync.WaitGroup{}
	group.Add(2)
	go WaitReleaseCaller("a", group)
	go WaitReleaseCaller("b", group)
	group.Wait()

	// normal way (non-blocking)
	go NormalCaller("c")
}
