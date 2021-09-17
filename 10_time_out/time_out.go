package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("num = ", num)
			case <-time.After(time.Second * 3):
				fmt.Println("超时")
				ch2 <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch1 <- i
		time.Sleep(time.Second)
	}

	b := <-ch2

	fmt.Println(b)
	fmt.Println("程序结束")

}
