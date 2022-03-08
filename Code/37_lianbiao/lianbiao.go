package main

import "fmt"

type ring struct {
	Val  int
	next *ring
}

func main() {
	head := &ring{}
	tail := head

	for i:=1;i<=5;i++  {
		tail.next = &ring{i,nil}
		tail = tail.next
	}

	head = head.next
	tail.next = head


	res := []int{}


	for head.next.next != head {
		res = append(res,head.next.Val)
		head.next = head.next.next
		head = head.next
	}

	fmt.Println(res)
	for i:=0;i<2;i++ {
		fmt.Printf("%d ",head.Val)
		head = head.next
	}
}
