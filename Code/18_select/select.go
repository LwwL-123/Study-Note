package main

import (
	"fmt"
	"time"
)

func addNumberToChan1(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(4 * time.Second)
	}
}

func addNumberToChan2(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(5 * time.Second)
	}
}

func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	go addNumberToChan1(chan1)
	go addNumberToChan2(chan2)

	for {
		select {
		case e := <- chan1 :
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <- chan2 :
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}