package main

import "time"

func main() {
	ch := make(chan int,5)
	ch <- 5

	close(ch)

	a := <- ch
	println(a)

	b := <-  ch
	println(b)

	go func() {
		defer recover()
		ch <-100

	}()


	time.Sleep(1 * time.Second)
}
