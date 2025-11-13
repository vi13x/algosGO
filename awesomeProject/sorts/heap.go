package sorts

// HeapSort выполняет пирамидальную сортировку и возвращает отсортированную копию слайса.
func HeapSort(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)

	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, heapSize, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, heapSize, largest)
	}
}
