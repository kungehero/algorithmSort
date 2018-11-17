package main

func ShellSort(arr []int) {
	length := len(arr)
	gap := len(arr) / 2
	for gap > 0 {
		for index := gap; index < length; index++ {
			j := index
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap

			}
		}
		gap /= 2
	}
}
