package main

import (
	"fmt"
)

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test_2_wei_bag_problem1(weight, value, 4)

}

func test_2_wei_bag_problem1(weight, value []int, cap int) {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	// 初始化
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = 0
		}
	}

	for j := weight[0]; j < len(dp[0]); j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < len(dp); i++ {
		for j := weight[i]; j < len(dp[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	fmt.Println(dp)

}

func max(a ...int) int {
	res := -1
	for _,j := range a {
		if j > res{
			res = j
		}
	}

	return res
}
