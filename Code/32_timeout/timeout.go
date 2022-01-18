package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	fmt.Println(time.Now())
	go func() {
		time.Sleep(500 * time.Millisecond)
		c <- 0
		time.Sleep(500 * time.Millisecond)
		c <- 1
	}()

	for {
		select {
		case p := <-c:
			fmt.Printf("p=%d\n", p)
		case <-time.After(1 * time.Second):
			fmt.Println(time.Now())
			fmt.Printf("timeout")
			return
		}
	}
}
