package sorts

// MergeSort реализует сортировку слиянием.
func MergeSort(arr []int, cb StepCallback) {
	if len(arr) <= 1 {
		return
	}
	buffer := make([]int, len(arr))
	mergeSort(arr, buffer, 0, len(arr)-1, cb)
}

func mergeSort(arr, buffer []int, left, right int, cb StepCallback) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(arr, buffer, left, mid, cb)
	mergeSort(arr, buffer, mid+1, right, cb)
	merge(arr, buffer, left, mid, right, cb)
}

func merge(arr, buffer []int, left, mid, right int, cb StepCallback) {
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			buffer[k] = arr[i]
			i++
		} else {
			buffer[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		buffer[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		buffer[k] = arr[j]
		j++
		k++
	}
	for idx := left; idx <= right; idx++ {
		arr[idx] = buffer[idx]
		emitStep(cb, arr)
	}
}
