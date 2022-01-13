package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	tmp := float64(6.66)
	fmt.Println(tmp)

	//tmpPtr := unsafe.Pointer(&tmp)
	//fmt.Println(tmpPtr)
	//
	//tmpP := (*uint64)(tmpPtr)
	//fmt.Println(tmpP)
	//
	//fmt.Println(*tmpP)

	res := *(*float32)(unsafe.Pointer(&tmp))

	fmt.Println(res)
}
