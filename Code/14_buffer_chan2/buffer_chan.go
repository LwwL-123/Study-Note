package main

import "fmt"

func main() {
	done := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("tread")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
