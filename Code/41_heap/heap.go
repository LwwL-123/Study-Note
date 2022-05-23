package main

import (
	"container/heap"
	"fmt"
)

type initHeap struct {
	heap []int
	// true大根堆 false小根堆
	bool
}
func main() {
	min := &initHeap{[]int{3,6,2,8,5},false}
	//max := &initHeap{[]int{},true}
	heap.Init(min)
	//heap.Init(max)
	//for {
	//	var tmp int
	//	fmt.Scan(&tmp)
	//	insert(min,max,tmp)
	//	fmt.Println(max)
	//	fmt.Println(min)
	//	fmt.Println(getM(min,max,tmp))
	//}
	fmt.Println(min)
}

func getM(min,max *initHeap,i int) int {
	if max.Len() == min.Len() {
		return (min.heap[0] + max.heap[0]) /2
	}else {
		return max.heap[0]
	}
}

func insert(min,max *initHeap,i int) {
	if min.Len() == 0 && max.Len() == 0 {
		heap.Push(max,i)
	}else if min.Len() == max.Len() {// 如果大根堆小根堆长度相等，则必须要插入大根堆
		if i > min.heap[0] {
			heap.Push(max,heap.Pop(min))
			heap.Push(min,i)
		}else {
			heap.Push(max,i)
		}
	}else {
		if i < max.heap[0] {
			heap.Push(min,heap.Pop(max))
			heap.Push(max,i)
		}else {
			heap.Push(min,i)
		}
	}


}


func (h initHeap) Len() int{ return len(h.heap) }
func (h initHeap) Less(i, j int) bool {
	if h.bool {
		return h.heap[i] > h.heap[j]
	}else {
		return h.heap[i] < h.heap[j]
	}
}
func (h initHeap) Swap(i, j int) { h.heap[i],h.heap[j] = h.heap[j],h.heap[i] }
func (h *initHeap) Push(x interface{}) { h.heap = append(h.heap, x.(int)) }
func (h *initHeap) Pop() interface{} {
	old := *h
	l := len(old.heap)
	x := old.heap[l-1]
	h.heap = old.heap[:l-1]
	return x
}


// 最大值
func max



