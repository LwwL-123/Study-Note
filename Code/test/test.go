package main

import (
	"fmt"
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


		fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))
		fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))
		fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))
		fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))
		fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))
		fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYCJY"))
		fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{}))

}