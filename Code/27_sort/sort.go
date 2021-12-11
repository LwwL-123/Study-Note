package main

import (
	"fmt"
	"sort"
)

func main()  {

	a := [][]int{{2,4},{2,5},{3,6},{1,5}}

	sort.Slice(a, func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		}else if a[i][0] == a[j][0]{
			if a[i][1] > a[j][1] {
				return true
			}else {
				return false
			}
		}else {
			return false
		}

	})

	var height []int
	for i := range a {
		height = append(height,a[i][1])
	}

	fmt.Println(a)
	fmt.Println(height)
}