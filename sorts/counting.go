package sorts

// CountingSort реализует сортировку подсчётом и поддерживает отрицательные значения.
func CountingSort(arr []int, cb StepCallback) {
	if len(arr) == 0 {
		return
	}
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	rangeVal := maxVal - minVal + 1
	count := make([]int, rangeVal)
	output := make([]int, len(arr))

	for _, v := range arr {
		count[v-minVal]++
	}
	for i := 1; i < rangeVal; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		count[val-minVal]--
		output[count[val-minVal]] = val
	}
	copy(arr, output)
	emitStep(cb, arr)
}
