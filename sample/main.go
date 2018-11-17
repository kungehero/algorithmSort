package main

import (
	_ "database/sql/driver"
	"fmt"

	_ "database/sql"

	_ "github.com/lib/pq"
)

//交换  冒泡排序   快速排序  鸡尾酒
//插入   二分插入（折半查找  插入排序） 希尔排序
//选择  直接选择  堆排序

func main() {
	arr := []int{4, 6, 7, 2, 1, 3, 9}
	Heap(arr)

	fmt.Println(arr)

}
func Heap(arr []int) {
	for index := 0; index < len(arr); index++ {
		AdjustHeap(arr[:len(arr)-index])
		arr[0], arr[len(arr)-index-1] = arr[len(arr)-index-1], arr[0]
	}

}
func AdjustHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		if arr[2*i+1] > arr[i] {
			arr[2*i+1], arr[i] = arr[i], arr[2*i+1]
		}
		if 2*i+2 < len(arr) {
			if arr[2*i+2] < arr[i] {
				arr[2*i+1], arr[i] = arr[i], arr[2*i+1]
			}
		}
	}
	for i := 0; i <= len(arr)/2-1; i++ {
		if arr[2*i+1] > arr[i] {
			arr[2*i+1], arr[i] = arr[i], arr[2*i+1]
		}
		if 2*i+2 < len(arr) {
			if arr[2*i+2] < arr[i] {
				arr[2*i+1], arr[i] = arr[i], arr[2*i+1]
			}
		}
	}
}
