package algorithmSort

import (
	"container/list"
)

func bubbleSort(arr []int) {
	L := len(arr)
	for i := 0; i < L; i++ {
		for j := 0; j < L-i; j++ {
			Swap(j-1, j, arr)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreTraversal(node *TreeNode) []interface{} {
	arr := make([]interface{}, 0)
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		arr = append(arr, node.Val)
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			node := v.Value.(*TreeNode)
			node = node.Right
			stack.Remove(v)
		}
	}
	return arr
}

func Swap(i, j int, arr []int) {
	if arr[i] > arr[j] {
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
	}
}
