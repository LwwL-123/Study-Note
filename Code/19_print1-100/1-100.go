package main

import (
	"fmt"
	"sync"
)

var ch1 = make(chan struct{}, 0)
var ch2 = make(chan struct{}, 0)
var ch3 = make(chan struct{}, 0)
var wg sync.WaitGroup

func main() {
	wg.Add(1)

	go a()
	go b()
	go c()

	ch1 <- struct{}{}
	wg.Wait()
}

func a() {
	for {
		select {
		case <-ch1:
			for i := 0; i < 100; i++ {
				fmt.Printf("我是第%d个1\n", i+1)
			}
			ch2 <- struct{}{}
			return
		}

	}
}
func b() {
	for {
		select {
		case <-ch2:
			for i := 0; i < 100; i++ {
				fmt.Printf("我是第%d个2\n", i+1)
			}
			ch3 <- struct{}{}
			return
		}

	}

}
func c() {
	for {
		select {
		case <-ch3:
			for i := 0; i < 100; i++ {
				fmt.Printf("我是第%d个3\n", i+1)
			}
			wg.Done()

			return
		}

	}
}
