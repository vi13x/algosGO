package sorts

// RadixSort выполняет поразрядную сортировку и возвращает отсортированную копию слайса.
func RadixSort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	negatives := make([]int, 0, len(input))
	nonNegatives := make([]int, 0, len(input))
	for _, v := range input {
		if v < 0 {
			negatives = append(negatives, -v)
		} else {
			nonNegatives = append(nonNegatives, v)
		}
	}

	radixLSD(nonNegatives)
	radixLSD(negatives)

	result := make([]int, 0, len(input))
	for i := len(negatives) - 1; i >= 0; i-- {
		result = append(result, -negatives[i])
	}
	result = append(result, nonNegatives...)
	return result
}

func radixLSD(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	exp := 1
	buffer := make([]int, len(arr))
	for maxVal/exp > 0 {
		count := make([]int, 10)
		for _, v := range arr {
			digit := (v / exp) % 10
			count[digit]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			digit := (arr[i] / exp) % 10
			count[digit]--
			buffer[count[digit]] = arr[i]
		}
		copy(arr, buffer)
		exp *= 10
	}
}
