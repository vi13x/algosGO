package sorts

// RadixSort реализует поразрядную сортировку для неотрицательных чисел.
func RadixSort(arr []int, cb StepCallback) {
	if len(arr) == 0 {
		return
	}
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	exp := 1
	output := make([]int, len(arr))
	for maxVal/exp > 0 {
		count := make([]int, 10)
		for i := 0; i < len(arr); i++ {
			digit := (arr[i] / exp) % 10
			count[digit]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			digit := (arr[i] / exp) % 10
			count[digit]--
			output[count[digit]] = arr[i]
		}
		copy(arr, output)
		emitStep(cb, arr)
		exp *= 10
	}
}
