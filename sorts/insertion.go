package sorts

// InsertionSort сортирует массив методом вставок.
func InsertionSort(arr []int, cb StepCallback) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			emitStep(cb, arr)
		}
		arr[j+1] = key
		emitStep(cb, arr)
	}
}
