package main

func main() {
	go InfiniteLoopSleep1Sec("a")
	// in below call, adding 'go' prefix does not impact output
	// because despite go primitive, select block after it implicitly blocks execution execution
	// making it pretty much equivalent to conventional (non-goroutine) invocation
	InfiniteLoopSleep2Sec("b")
	select {}
}
