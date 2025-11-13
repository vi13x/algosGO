package sorts

// QuickSort реализует быструю сортировку.
func QuickSort(arr []int, cb StepCallback) {
	quickSort(arr, 0, len(arr)-1, cb)
}

func quickSort(arr []int, low, high int, cb StepCallback) {
	if low >= high {
		return
	}
	p := partition(arr, low, high, cb)
	quickSort(arr, low, p-1, cb)
	quickSort(arr, p+1, high, cb)
}

func partition(arr []int, low, high int, cb StepCallback) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			emitStep(cb, arr)
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	emitStep(cb, arr)
	return i
}
