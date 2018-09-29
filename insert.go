package algorithmSort

import (
	"container/list"
)

func InsertSort(arr []int) {
	L := len(arr)
	for i := 1; i < L; i++ {
		for j := i; j > 0; j-- {
			Swap(j-1, j, arr)
		}
	}
}

func InOrderTraversal(node *TreeNode) []interface{} {
	arr := make([]interface{}, 0)
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			node := v.Value.(*TreeNode)
			arr = append(arr, node.Val)
			node = node.Right
			stack.Remove(v)
		}
	}
	return arr
}
