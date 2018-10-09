package main

import "fmt"

//交换  冒泡排序   快速排序  鸡尾酒
//插入   二分插入（折半查找  插入排序） 希尔排序
//选择  直接选择  堆排序
func main() {
	arr := []int{2, 1, 3, 5, 4, 7, 6}
	Select(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func insert(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func Select(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			temp := a[i]
			a[i] = a[min]
			a[min] = temp
		}
	}
}
