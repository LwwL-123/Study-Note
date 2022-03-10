package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {

	l1 := list.New()
	l2 := list.New()
	var num1,num2 int
	fmt.Scan(&num1)

	for i:=0;i<num1;i++{
		var tmp int
		fmt.Scan(&tmp)
		l1.PushBack(tmp)
	}

	fmt.Scan(&num2)
	for i:=0;i<num2;i++{
		var tmp int
		fmt.Scan(&tmp)
		l2.PushBack(tmp)
	}


	l3 := list.New()
	res := l3.Back()

	e1 := l1.Back()
	e2 := l2.Back()
	for e1 != nil && e2 != nil {
		val := e1.Value.(int) + e2.Value.(int)

		if val / 10 > 0 {
			if res != nil {
				res.Value = res.Value.(int) + (val % 10)
			}else {
				l3.PushFront(val % 10)
			}

			l3.PushFront(1)
			res = l3.Front()
		}else {
			if res != nil {
				res.Value = res.Value.(int) + (val % 10)
				res = l3.Front().Prev()
			}else {
				l3.PushFront(val)
				res = l3.Front().Prev()
			}
		}

		e1 = e1.Prev()
		e2 = e2.Prev()
	}

	for e1 != nil {
		l3.PushFront(e1.Value)
		e1 = e1.Prev()
	}

	for e2 != nil {
		l2.PushFront(e2.Value)
		e2 = e2.Prev()
	}

	print(l3)

}
func print(l1 *list.List) {
	for l := l1.Front();l!=nil;l=l.Next() {
		fmt.Printf("%d ", l.Value)
	}
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


	a.Do(func(i interface{}) {
		fmt.Println(i.(string))
	})

}