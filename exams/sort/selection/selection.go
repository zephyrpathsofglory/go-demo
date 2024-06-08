package selection

func SelectionSort(arr []int) {
	lenz := len(arr)
	for i := 0; i < lenz; i++ {
		for j := i + 1; j < lenz; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
