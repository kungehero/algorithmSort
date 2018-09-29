package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{Val: 1}
	node.Right = &TreeNode{Val: 2}
	node.Right.Left = &TreeNode{Val: 3}

	for _, v := range PostOrderSort(node) {
		fmt.Println(v)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
