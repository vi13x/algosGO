package sorts

const minRun = 32

// TimSort выполняет Timsort и возвращает отсортированную копию слайса.
func TimSort(input []int) []int {
	if len(input) <= 1 {
		return append([]int(nil), input...)
	}

	arr := make([]int, len(input))
	copy(arr, input)

	n := len(arr)
	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		insertionSortRange(arr, start, end)
	}

	for size := minRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size
			right := left + 2*size
			if mid > n {
				mid = n
			}
			if right > n {
				right = n
			}
			if mid < right {
				merge(arr, make([]int, len(arr)), left, mid, right)
			}
		}
	}

	return arr
}

func insertionSortRange(arr []int, start, end int) {
	for i := start + 1; i < end; i++ {
		key := arr[i]
		j := i - 1
		for j >= start && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
