package main

import "fmt"

type t struct {
	name string
	age  *Age
}

type Age struct {
	a int
}

func main() {
	r := getT(1)

	if r.age.a == 10 {
		fmt.Println("yes")
	}
}

func getT(num int) (reply *t) {
	tmp := &t{name: "111", age: &Age{a: 10}}

	if num == 1 {
		return
	}

	reply = tmp
	return
}
