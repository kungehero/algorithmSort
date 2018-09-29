package main

import (
	"fmt"

	"github.com/kungehero/algorithmSort"
)

func main() {
	node := &algorithmSort.TreeNode{Val: 1}
	node.Right = &algorithmSort.TreeNode{Val: 2}
	node.Right.Left = &algorithmSort.TreeNode{Val: 3}
	arr := algorithmSort.PostOrderSort(node)
	for _, v := range arr {
		fmt.Println(v)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
