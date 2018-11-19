func SecSort(arr []int) []int {
	if len(arr) > 1 {
		key := len(arr) / 2
		left := SecSort(arr[:key])
		right := SecSort(arr[key:])
		return MergeSort(left, right)
	} else {
		return arr
	}
}
func MergeSort(a, b []int) []int {
	i, j := 0, 0
	var k []int
	for i < len(a) && j < len(b) {

		if a[i] < b[j] {
			k = append(k, a[i])
			i++
		} else {
			k = append(k, b[j])
			j++
		}
	}
	for len(a)-i > 0 {
		k = append(k, a[i])
		i++
	}
	for len(b)-j > 0 {
		k = append(k, b[j])
		j++
	}

	return k
}
