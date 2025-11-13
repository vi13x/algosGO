package sorts

// HeapSort реализует пирамидальную сортировку.
func HeapSort(arr []int, cb StepCallback) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, cb)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		emitStep(cb, arr)
		heapify(arr, i, 0, cb)
	}
}

func heapify(arr []int, n, i int, cb StepCallback) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		emitStep(cb, arr)
		heapify(arr, n, largest, cb)
	}
}
