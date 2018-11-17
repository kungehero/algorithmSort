package algorithmSort

func QuickSort(a []int, left, right int) {
	//从右向左找出第一个小于key(基准数)的值
	//从左到右找出第一个大于key的值
	//当左右移动相等时 去除基准值的索引

	//46721395
	if left >= right {
		return
	}
	i, j, key := left, right, a[left]
	for i < j {
		for i < j && a[j] > key {
			j--
		} //- j=3
		if i < j {
			a[i], a[j] = a[j], a[i]
			i++
		}
		//67241395
		for i < j && a[i] < key {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			j--
		}
	}
	a[i] = key
	QuickSort(a, left, i-1)
	QuickSort(a, i+1, right)
}
