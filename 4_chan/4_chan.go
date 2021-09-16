package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(3)
	go incCount()
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(counter)
}

func incCount() {
	defer wg.Done()
	//for i := 0; i < 2; i++ {
	//	value := counter
	//	runtime.Gosched()
	//	value++
	//	counter = value
	//}
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加1
		runtime.Gosched()
	}
}
