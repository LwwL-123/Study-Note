package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println(counter)
}

func incCounter() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}

}
