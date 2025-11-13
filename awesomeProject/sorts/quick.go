package sorts

// QuickSort выполняет быструю сортировку и возвращает отсортированную копию слайса.
func QuickSort(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quickSort(arr, low, pivotIndex-1)
	quickSort(arr, pivotIndex+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
