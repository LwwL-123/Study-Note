package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)


func main()  {

	// 创建分析文件
	file, err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Printf("创建采集文件失败, err:%v\n", err)
		return
	}

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()


	nums := []int{1,5,1}
	NextPermutation(nums)
	fmt.Println(nums)

	time.Sleep(10 * time.Second)

}

func NextPermutation(nums []int)  {

	i := len(nums) - 2
	// 从右往左遍历，得到第一个左边小于右边的数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果i不是最后一个序列
	if i >= 0 {
		j := len(nums) - 1
		// 找到从右往左第一个比A[i]大的数
		for nums[j] <= nums[i] {
			j--
		}
		// 交换两个数
		nums[i],nums[j] = nums[j],nums[i]
	}

	// 交换A[j:end]
	// 因为i右边的数是从右往左递增的，交换ij后，仍然保持单调递增特性
	// 此时需要从头到尾交换
	for l,r := i + 1, len(nums) - 1; l < r; l,r = l+1,r-1 {
		nums[l],nums[r] = nums[r],nums[l]
	}
}