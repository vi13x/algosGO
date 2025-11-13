package sorts

const minRun = 32

// TimSort реализует упрощённый вариант Timsort.
func TimSort(arr []int, cb StepCallback) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for start := 0; start < n; start += minRun {
		end := start + minRun - 1
		if end >= n {
			end = n - 1
		}
		insertionSortRange(arr, start, end, cb)
	}

	size := minRun
	for size < n {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			right := left + 2*size - 1
			if mid >= n-1 {
				break
			}
			if right >= n {
				right = n - 1
			}
			mergeRange(arr, left, mid, right, cb)
		}
		size *= 2
	}
}

func insertionSortRange(arr []int, left, right int, cb StepCallback) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			emitStep(cb, arr)
		}
		arr[j+1] = key
		emitStep(cb, arr)
	}
}

func mergeRange(arr []int, left, mid, right int, cb StepCallback) {
	n1 := mid - left + 1
	n2 := right - mid
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)
	copy(leftArr, arr[left:left+n1])
	copy(rightArr, arr[mid+1:mid+1+n2])

	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
		emitStep(cb, arr)
	}
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
		emitStep(cb, arr)
	}
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
		emitStep(cb, arr)
	}
}
