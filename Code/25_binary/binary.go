package main

func main()  {
	nums := []int{1,2,3,4,5,28,5689}
	res := binary(nums,5689)
	println(res)

}

func binary(nums []int,target int) int{

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}