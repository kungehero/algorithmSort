package main

import "fmt"

//交换  冒泡排序   快速排序  鸡尾酒
//插入   二分插入（折半查找  插入排序） 希尔排序
//选择  直接选择  堆排序
func main() {
	arr := []int{2, 1, 3, 5, 4, 7, 6}
	Partion(arr, 0, len(arr)-1)
	for _, v := range arr {
		fmt.Println(v)
	}
}
