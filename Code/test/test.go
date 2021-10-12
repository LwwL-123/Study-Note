package main

import "fmt"

func main() {
	aaa := []int {100,200,300,400,500}
	for i,x := range aaa{
		fmt.Println(i)
		fmt.Println(x)
	}
}
