package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//h := ConstructorHeap([]int{1, 5, 4, 3, 7, 2, 1, 6, 9, 9, 9}, 3)
	//fmt.Println(h.data)
	minHeap := &hp{isBigHeap: false}
	maxHeap := &hp{isBigHeap: true}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	nums := []int{3, 6, 2, 8, 5}
	for _, num := range nums {
		heap.Push(minHeap, num)
	}

	a := heap.Pop(minHeap)
	fmt.Println(a)
	fmt.Println(minHeap)
}

type minHeap struct {
	len  int
	data []int
}

func ConstructorHeap(nums []int, k int) *minHeap {
	h := &minHeap{
		len:  k,
		data: nil,
	}

	for _, num := range nums {
		h.add(num)
	}

	return h
}

func (this *minHeap) add(num int) {
	if len(this.data) < this.len {
		this.data = append(this.data, num)
		this.up(len(this.data) - 1)
	} else if num > this.data[0] {
		this.data[0] = num
		this.down(0)
	}
}

func (this *minHeap) up(index int) {
	for index > 0 {
		pre := (index - 1) / 2
		if this.data[index] < this.data[pre] {
			this.data[index], this.data[pre] = this.data[pre], this.data[index]
			index = pre
		} else {
			break
		}
	}
}

func (this *minHeap) down(index int) {
	for index*2+1 < len(this.data) {
		// 左孩子
		child := index*2 + 1
		// 是否存在右孩子,并将child指向两个孩子中最小的那个
		if child+1 < len(this.data) && this.data[child+1] < this.data[child] {
			child++
		}
		if this.data[index] > this.data[child] {
			this.data[index], this.data[child] = this.data[child], this.data[index]
			index = child
		} else {
			break
		}
	}
}
