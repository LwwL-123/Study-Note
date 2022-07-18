package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nodeVal := []int{1, 2, 3, 4, 5, 6, 7}
	root := constructBinaryTree(nodeVal)
	printBinaryTree(root, len(nodeVal))
	spew.Dump(root)
}
func constructBinaryTree(array []int) *TreeNode {
	var root *TreeNode
	// 初始化所有node节点
	nodes := make([]*TreeNode, len(array))
	for i := 0; i < len(nodes); i++ {
		if array[i] != -1 {
			nodes[i] = &TreeNode{array[i], nil, nil}
		}
		if i == 0 {
			root = nodes[i]
		}
	}

	// 连接所有节点
	for i := 0; i*2+2 < len(array); i++ {
		if nodes[i] != nil {
			nodes[i].Left = nodes[2*i+1]
			nodes[i].Right = nodes[2*i+2]
		}

	}

	return root
}

func printBinaryTree(root *TreeNode, n int) {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	result := []int{}
	for len(queue) > 0 {
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			if node != nil {
				result = append(result, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			} else {
				result = append(result, -1)
			}
		}
		// 清除队列中的本层节点, 进入下一层遍历
		queue = queue[len(queue):]
	}

	// 参数n控制输出值数量, 否则二叉树最后一层叶子节点的孩子节点也会被打印(但是这些孩子节点是不存在的).
	fmt.Println(result[:n])
}
