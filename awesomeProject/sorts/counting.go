package sorts

// CountingSort выполняет сортировку подсчётом и возвращает отсортированную копию слайса.
func CountingSort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	arr := make([]int, len(input))
	copy(arr, input)

	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	rangeSize := maxVal - minVal + 1
	counts := make([]int, rangeSize)
	for _, v := range arr {
		counts[v-minVal]++
	}

	index := 0
	for valueOffset, count := range counts {
		for count > 0 {
			arr[index] = valueOffset + minVal
			index++
			count--
		}
	}

	return arr
}
