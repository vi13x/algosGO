package sorts

// MergeSort выполняет сортировку слиянием и возвращает отсортированную копию слайса.
func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return append([]int(nil), input...)
	}

	arr := make([]int, len(input))
	copy(arr, input)
	buffer := make([]int, len(arr))
	mergeSort(arr, buffer, 0, len(arr))
	return arr
}

func mergeSort(arr, buffer []int, start, end int) {
	if end-start <= 1 {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, buffer, start, mid)
	mergeSort(arr, buffer, mid, end)
	merge(arr, buffer, start, mid, end)
}

func merge(arr, buffer []int, start, mid, end int) {
	i, j, k := start, mid, start
	for i < mid && j < end {
		if arr[i] <= arr[j] {
			buffer[k] = arr[i]
			i++
		} else {
			buffer[k] = arr[j]
			j++
		}
		k++
	}
	for i < mid {
		buffer[k] = arr[i]
		i++
		k++
	}
	for j < end {
		buffer[k] = arr[j]
		j++
		k++
	}
	for idx := start; idx < end; idx++ {
		arr[idx] = buffer[idx]
	}
}
