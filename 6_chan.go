package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("start")
		ch <- 0
		fmt.Println("exit")
	}()
	fmt.Println("wait goroutine")

	<-ch
	fmt.Println("wait goroutine")

}
