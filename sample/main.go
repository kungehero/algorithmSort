package main

import "fmt"

//交换  冒泡排序   快速排序  鸡尾酒
//插入   二分插入（折半查找  插入排序） 希尔排序
//选择  直接选择  堆排序
func main() {
	arr := []interface{}{2, 1, 3, 5, 4, 7, 6}
	headSoft(arr, func(a, b interface{}) bool {
		a1 := a.(int)
		a2 := b.(int)
		return a1 <= a2
	})
	fmt.Println(arr)

	for _, v := range arr {
		fmt.Println(v)
	}
}
func headSoft(array []interface{}, cmp func(a, b interface{}) bool) {
	len := len(array)
	//建堆
	for i := (len / 2) - 1; i >= 0; i-- {
		adjustHeap(array, i, len, cmp)
	}

	//排序
	for i := 0; i < len; i++ {
		sz := len - i - 1
		array[0], array[sz] = array[sz], array[0]
		adjustHeap(array[0:sz], 0, sz, cmp)
	}
}

func adjustHeap(array []interface{}, i, len int, cmp func(a, b interface{}) bool) {
	for ch := i<<1 + 1; ch < len; ch = i<<1 + 1 {
		rg := i<<1 + 2
		if rg < len && cmp(array[ch], array[rg]) {
			ch = rg
		}
		if cmp(array[ch], array[i]) {
			return
		}
		array[ch], array[i] = array[i], array[ch]
		i = ch
	}
}
