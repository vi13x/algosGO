package sorts

// InsertionSort выполняет сортировку вставками и возвращает отсортированную копию слайса.
func InsertionSort(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}
