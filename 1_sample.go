package main

import "fmt"

func main() {

	var n [10]int
	for i := 0; i < 10; i++ {
		n[i] = i + 100
		fmt.Println(n[i])
	}

	//指针
	var ip *int
	a := 10
	ip = &a
	fmt.Printf("a 的值：%d\n", a)
	fmt.Printf("a 的地址：%x\n", &a)
	fmt.Printf("ip 的地址：%x\n", ip)
	fmt.Printf("ip 的值：%d\n", *ip)

	*ip = 100

	fmt.Printf("a 的值：%d\n", a)
	fmt.Printf("a 的地址：%x\n", &a)
	fmt.Printf("ip 的地址：%x\n", ip)
	fmt.Printf("ip 的值：%d\n", *ip)
}
