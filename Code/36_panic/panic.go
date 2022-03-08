package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("111")
		panic("555")
		fmt.Println("222")
	}()

	defer func() {
		fmt.Println("777")
		panic("666")
		fmt.Println("666")
	}()

	fmt.Println("333")
	panic("panic")

	fmt.Println("444")

}
