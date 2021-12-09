package main

func main()  {
	
}

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := leftSearch(nums, target)
	r := rightSearch(nums, target)

	return []int{l, r}
}

func leftSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// 搜索区间为[left,right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 区间变为[mid+1,right]
			left = mid + 1
		} else if nums[mid] > target {
			// 区间变为[left,mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩右侧边界
			right = mid - 1
		}
	}

	// 因为退出条件为left == right + 1 ,所以当target比nums中所有值都大时,right为最右侧的值,left=right+1,产生越界
	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func rightSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// 搜索区间为[left,right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 区间变为[mid+1,right]
			left = mid + 1
		} else if nums[mid] > target {
			// 区间变为[left,mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩左侧边界
			left = mid + 1
		}
	}

	// 因为退出条件为left == right + 1 ,所以当target比nums中所有值都小时,mid为0,right=mid-1,产生越界
	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}
