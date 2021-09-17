package main

import "fmt"

func main() {
	//创建一个缓冲为3的通道
	ch := make(
		chan int,
		3,
	)
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch))
}
