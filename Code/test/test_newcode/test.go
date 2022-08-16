package main

import (
	"fmt"
)

func main() {
	a := firstMissingPositive([]int{7, 8, 9, 11, 12})
	fmt.Println(a)
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for k, _ := range nums {
		if nums[k] != k+1 {
			return k + 2
		}
	}
	return len(nums) + 1
}
