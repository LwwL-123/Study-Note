package main

import (
	"container/heap"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	minHeap := &hp{isBigHeap: false}
	maxHeap := &hp{isBigHeap: true}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	nums1 = append(nums1, nums2...)
	for _, num := range nums1 {
		insert(minHeap, maxHeap, num)
	}

	return getMid(minHeap, maxHeap)

}

func getMid(min, max *hp) float64{
	if min.Len() == max.Len() {
		return float64(min.IntSlice[0] + max.IntSlice[0])/2)
	}else {
		return float64(max.IntSlice[0])
	}

	return 0
}

func insert(min, max *hp, num int) {
	if min.Len() == 0 && max.Len() == 0 {
		heap.Push(max, num)
	} else if min.Len() == max.Len() {
		if num > max.IntSlice[0] {
			heap.Push(max, heap.Pop(min))
			heap.Push(min, num)
		} else {
			heap.Push(max, num)
		}
	} else {
		if num < max.IntSlice[0] {
			heap.Push(min, heap.Pop(max))
			heap.Push(max, num)
		} else {
			heap.Push(min, num)
			sort.Slice(min.IntSlice, func(i, j int) bool {
				return min.IntSlice[i] < min.IntSlice[j]
			}
		}
	}
}

type hp struct {
	sort.IntSlice
	isBigHeap bool
}

func (h hp) Less(i, j int) bool {
	if h.isBigHeap {
		return h.IntSlice[i] > h.IntSlice[j]
	} else {
		return h.IntSlice[i] < h.IntSlice[j]
	}
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	//a := h.IntSlice
	//v := a[len(a)-1]
	//h.IntSlice = a[:len(a)-1] // replicate 0 to (len(a)-1)(exclude)
	//return v

	res := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return res
}
