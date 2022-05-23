package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{7, 5, 3, 6, 4, 2, 1}
	fmt.Println(minimum(a))

}



/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型一维数组
 * @return long长整型
 */
func minimum( a []int ) int64 {
	length := len(a)
	res := math.MaxInt64
	a = append(a, a...)
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			lSum,rSum := sum(a[i:i+j]),sum(a[i+j:i+length])
			res = int(min(int64(res), abs(lSum, rSum)))
		}
	}

	return int64(res)
}

func min(a,b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func sum(nums []int) int64 {
	var res int64
	for _,j := range nums {
		res += int64(j)
	}
	return res
}

func abs(a,b int64) int64 {
	if a - b > 0 {
		return a - b
	}

	return b - a
}