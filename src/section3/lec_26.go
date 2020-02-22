package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MIN_DELAY_MILLIS = 800
const MAX_DELAY_MILLIS = 1000
const ACK_DELAY_MILLIS = 100

func PingPong(name string, rxChannel chan int, txChannel chan int) {
	for rxBytes := range rxChannel {
		if name == "ponger" {
			name = "  ponger"
		}

		// receiving and processing
		time.Sleep(time.Millisecond * ACK_DELAY_MILLIS)
		fmt.Printf("%s <-- received %d bytes\n", name, rxBytes)
		time.Sleep(time.Millisecond * time.Duration(rxBytes))

		// transmitting
		var txBytes int = MIN_DELAY_MILLIS + rand.Intn((MAX_DELAY_MILLIS - MIN_DELAY_MILLIS))
		fmt.Printf("%s sent --> %d bytes\n", name, txBytes)
		time.Sleep(time.Millisecond * ACK_DELAY_MILLIS)
		txChannel <- txBytes
	}
}

func main() {
	var chan1 chan int = make(chan int, 1)
	var chan2 chan int = make(chan int, 1)

	go PingPong("pinger", chan1, chan2)
	go PingPong("ponger", chan2, chan1)

	chan1 <- MIN_DELAY_MILLIS + rand.Intn((MAX_DELAY_MILLIS - MIN_DELAY_MILLIS))

	select {}
}
