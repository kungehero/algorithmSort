package main

import (
	"fmt"
)

func main() {
	/*node := &algorithmSort.TreeNode{Val: 1}
	node.Right = &algorithmSort.TreeNode{Val: 2}
	node.Right.Left = &algorithmSort.TreeNode{Val: 3}
	arr := algorithmSort.PostOrderSort(node)*/
	arr := []int{2, 1, 3, 5, 4, 7, 6}
	SelectSort(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
func SelectSort(arr []int) {
	L := len(arr)
	for i := 0; i < L-1; i++ {
		min := i
		for j := i; j < L; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			Swap(i, min, arr)
		}
	}
}
func Swap(i, j int, arr []int) {
	if arr[i] > arr[j] {
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
	}
}
