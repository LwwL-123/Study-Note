package main

import (
	"fmt"
	"runtime"
	"time"
)

var JobQueue = make(chan func(), 100)

func main() {
	for i := 0; i < 100; i++ {
		go Worker()
	}
	// [1,2,3]    [1,4,7]
	// [4,5,6]	  [2,5,8]
	// [7,8,9]	  [3,6,9]
	nums := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}}
	// 对角线反转
	for tmp := 0; tmp < len(nums)/2; tmp++ {
		n := len(nums)
		i := tmp
		JobQueue <- func() {
			time.Sleep(time.Duration((i+1)*5) * time.Second)
			fmt.Printf("对角线旋转第%d层", i)
			for j := 0; j < (n+1)/2; j++ {
				nums[i][j], nums[n-j-1][i], nums[n-i-1][n-j-1], nums[j][n-i-1] =
					nums[n-j-1][i], nums[n-i-1][n-j-1], nums[j][n-i-1], nums[i][j]
			}
		}
	}
	//// 中轴线反转
	//for i := 0; i < len(nums); i++ {
	//	fmt.Printf("注册第%d行", i)
	//	tmp := i
	//	JobQueue <- func() {
	//		time.Sleep(time.Duration((tmp+1)*5) * time.Second)
	//		left, right := 0, len(nums[tmp])-1
	//		for left < right {
	//			nums[tmp][left], nums[tmp][right] = nums[tmp][right], nums[tmp][left]
	//			left++
	//			right--
	//		}
	//		fmt.Println(tmp)
	//	}
	//}

	// 阻塞主线程
	for {
		fmt.Printf("协程数量为:%d\n", runtime.NumGoroutine())
		fmt.Println(nums)
		time.Sleep(1 * time.Second)
	}

}

func Worker() {
	for {
		select {
		case job := <-JobQueue:
			job()
		}
	}
}

//package main
//
//import "fmt"
//
//func main() {
//	ch := make(chan func(), 100)
//	for i := 0; i < 10; i++ {
//		tmp := i
//		ch <- func() {
//			fmt.Println(tmp)
//		}
//	}
//
//	for i := 0; i < 10; i++ {
//		a := <-ch
//		a()
//	}
//
//}
