package algorithmSort

import "fmt"

func ShellSort(arr []int) {

	gap := 2
	step := len(arr) / gap

	for step >= 1 {
		// 这里按步长开始每个分组的排序
		for i := step; i < len(arr); i++ {
			// 将按步长分组的子队列用直接插入排序算法进行排序
			InsertSort(arr, step)
		}
		// 完成一轮后再次缩小增量
		step /= gap

	}

	for _, v := range arr {
		fmt.Println(v)
	}
}

func InsertSort(tree []int, step int) {
	for i := step; i < len(tree); i++ {
		for j := i; j >= step && tree[j] < tree[j-step]; j -= step { //j=j-step
			tree[j], tree[j-step] = tree[j-step], tree[j]
		}
	}
}
