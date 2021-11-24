package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	i := 0
	j := 0
	go func() {
		for {
			if i < 5 {
				chan1 <- i
				i++
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if j < 5 {
				chan2 <- j
				j++
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case i := <-chan1:
			fmt.Println("chan1 ready.",i)
		case j := <-chan2:
			fmt.Println("chan2 ready.",j)
		}
		time.Sleep(5*time.Second)
	}

	println(math.Pow(2, 5))

	fmt.Println("main exit.")
}