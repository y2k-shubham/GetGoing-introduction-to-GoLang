package main

import "fmt"
import "time"

func HeavyFunc() {
	fmt.Println("HeavyFunc: Going-to-sleep")
	time.Sleep(time.Second * 2)
	fmt.Println("HeavyFunc: Woke-up")
}

func main() {
	fmt.Println("main: Invoking  HeavyFunc as GoRoutine")
	go HeavyFunc()
	fmt.Println("main: Completed HeavyFunc as GoRoutine")
	
	fmt.Println("main: Invoking  HeavyFunc normally")
	HeavyFunc()
	fmt.Println("main: Completed HeavyFunc normally")
}
