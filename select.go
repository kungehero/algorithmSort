package algorithmSort

import (
	"container/list"
)

func SelectSort(arr []int) {
	L := len(arr)
	for i := 1; i < L; i++ {
		min := i
		for j := i; j < L-1; j++ {
			if arr[j] < arr[j-1] {
				min = j
			}
		}
		if min != i {
			Swap(min, i, arr)
		}
	}
}

func PostOrderSort(node *TreeNode) []interface{} {
	arr := make([]interface{}, 0)
	var pre *TreeNode
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		v := stack.Back()
		top := v.Value.(*TreeNode)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && pre == top.Left) || pre == top.Right {
			arr = append(arr, top.Val)
			pre = top
			stack.Remove(v)

		} else {
			node = top.Right
		}
	}
	return arr
}
