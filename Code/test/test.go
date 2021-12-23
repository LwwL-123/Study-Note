package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	//cpunum := flag.Int("cpunum", 0, "cpunum")
	//flag.Parse()
	//fmt.Println("使用的cpunum:", *cpunum)
	//println("系统的CPU核数",runtime.NumCPU())
	//runtime.GOMAXPROCS(*cpunum)
	//for i := 0; i < *cpunum - 1; i++ {
	//	go func() {
	//		for {
	//
	//		}
	//	}()
	//}
	//for {
	//
	//}

	a := []int{-1,5,4,6,8,10,5,3,2,1}
	sort.Ints(a)
	fmt.Println(a)

	b := a[:1]
	c := a[:2]
	fmt.Println(b)
	fmt.Println(c)

	c[0] = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))
	fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))
	fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))
	fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYCJY"))
	fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{}))


}