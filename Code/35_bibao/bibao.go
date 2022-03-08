package main

import "fmt"

func main() {
	m := make(map[int]int,10)
	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	for k,v := range m {
		go func() {
			fmt.Println(k,v)
		}()
	}


	//for k,v := range m {
	//	go func(k,v int) {
	//		fmt.Println(k,v)
	//	}(k,v)
	//}

	for {}
}
