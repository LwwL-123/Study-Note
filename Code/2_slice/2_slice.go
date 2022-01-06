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

	//array := [10]int{1,2,3,4,5,6,7,8,9,10}
	//
	//var slice = array[0:5]
	//newSlice := slice[4:]
	//newSlice[0] = 100
	//
	//fmt.Println("lenth of slice: ", len(slice))
	//fmt.Println("capacity of slice: ", cap(slice))
	//fmt.Println(slice)
	//fmt.Println(newSlice)
	//fmt.Println(&slice[0] == &array[5])

	sliceA := make([]int, 5, 10)
	sliceB := sliceA[0:5]         //length = 5; capacity = 10
	sliceC := sliceA[2:]         //length = 8; capacity = 8



	fmt.Println(len(sliceB))
	fmt.Println(cap(sliceB))

	fmt.Println(len(sliceA))
	fmt.Println(cap(sliceA))

	fmt.Println(len(sliceC))
	fmt.Println(cap(sliceC))

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
