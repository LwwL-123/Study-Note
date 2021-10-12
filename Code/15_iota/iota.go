package main

import "fmt"

func main() {
	const (
		aaa = iota
		bbb
		ccc = iota + 1
		ddd = iota + 3
	)

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)
}
