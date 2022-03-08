package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	ring1()
}

func list1() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ring1() {
	a := ring.New(5)
	for i := 0; i < 5; i++ {
		a.Value = fmt.Sprintf("%d00",i)
		a = a.Next()
	}

	a.


	a.Do(func(i interface{}) {
		fmt.Println(i.(string))
	})



}