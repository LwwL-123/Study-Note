package main

import (
	"fmt"
)

func main() {
	//s := []int{1, 2, 3,3,3,3}
	//
	//for i,num := range s{
	//	if num == 3 {
	//		s = append(s[:i],s[i+1:]...)
	//		i--
	//
	//	}
	//}
	//
	//fmt.Println(s)
	//// 7为大小 9位最大容量
	//var ss = make([]int, 7, 9)
	//printSlice(s)
	//printSlice(ss)
	//
	//ss = append(ss, 10, 11, 12)
	//printSlice(ss)
	//
	//func(a int) { print(a + 1) }(5)
	//numbers := []int{0,1,2,3,4,5,6,7,8}
	///* 打印原始切片 */
	//fmt.Println("numbers ==", numbers)
	//
	///* 打印子切片从索引1(包含) 到索引4(不包含)*/
	//fmt.Println("numbers[1:4] ==", numbers[1:4])
	//
	///* 默认下限为 0*/
	//fmt.Println("numbers[:3] ==", numbers[:3])

	array := [10]int{1,2,3,4,5,6,7,8,9,10}

	var slice = array[0:10]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(slice)
	fmt.Println(&slice[0] == &array[5])
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
