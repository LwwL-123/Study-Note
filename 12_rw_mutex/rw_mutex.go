package main

import "sync"

var (
	count        int
	counterGuard sync.RWMutex
)

func getCount() int {
	counterGuard.Lock()

	defer counterGuard.Unlock()

	return count

}

func setCount(c int) {
	counterGuard.Lock()
	count = c
	counterGuard.Unlock()
}

func main() {
	setCount(1)
	a := getCount()
	println(a)
}
