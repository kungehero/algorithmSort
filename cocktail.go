package algorithmSort

func cocktail(src []int) {
	L := len(src)
	for i := 0; i < L/2; i++ {
		for j := i; j < L-i-1; j++ {
			if src[j] < src[j+1] {
				temp := src[j]
				src[j] = src[j+1]
				src[j+1] = temp
			}
		}
		//将最大值排到队头
		for j := L - 1 - (i + 1); j > i; j-- {
			if src[j] > src[j-1] {
				temp := src[j]
				src[j] = src[j-1]
				src[j-1] = temp
			}
		}
	}
}
