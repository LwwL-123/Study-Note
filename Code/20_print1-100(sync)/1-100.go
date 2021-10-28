package main

import (
	"sync"
)

var ab = make(chan struct{},0)
var bc = make(chan struct{},0)
var ca = make(chan struct{},0)
var wg sync.WaitGroup

func main()  {
	wg.Add(1)
	go a()
	go b()
	go c()

	ca <- struct{}{}
	wg.Wait()
}

func a()  {

	i:=0
	for {
		select {
		case <-ca:
			i++
			if i > 100 {
				wg.Done()
				return
			}
			print("1")
			ab <- struct{}{}
		}
	}
}

func b()  {

	for {
		select {
		case <-ab:
			print("2")
			bc <- struct{}{}
		}
	}
}

func c()  {
	for {
		select {
		case <-bc:
			print("3")
			ca <- struct{}{}
		}
	}
}