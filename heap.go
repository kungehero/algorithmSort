package algorithmSort

func HeadSort(array []interface{}, cmp func(a, b interface{}) bool) {
	len := len(array)
	//建堆
	for i := (len / 2) - 1; i >= 0; i-- {
		AdjustHeap(array, i, len, cmp)
	}

	//排序
	for i := 0; i < len; i++ {
		sz := len - i - 1
		array[0], array[sz] = array[sz], array[0]
		AdjustHeap(array[0:sz], 0, sz, cmp)
	}
}

func AdjustHeap(array []interface{}, i, len int, cmp func(a, b interface{}) bool) {
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
