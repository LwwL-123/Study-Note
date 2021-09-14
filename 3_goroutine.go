package main

import (
	"fmt"
	"time"
)

func main() {
	go DelayPrint()
	go HelloWorld()
	time.Sleep(2 * time.Second)
	//fmt.Println("main function")
}

func DelayPrint() {
	for i := 0; i <= 5; i++ {
		//time.Sleep(250*time.Millisecond)
		fmt.Println(i)
	}
}
func HelloWorld() {
	fmt.Println("Hello world goroutine")

}
