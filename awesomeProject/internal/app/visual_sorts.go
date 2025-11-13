package app

import (
	"math"
	"time"
)

func reportStep(data []int, delay time.Duration, report stepReporter) {
	snapshot := append([]int(nil), data...)
	report(snapshot)
	time.Sleep(delay)
}

func visualInsertionSort(data []int, delay time.Duration, report stepReporter) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
			reportStep(data, delay, report)
		}
		data[j+1] = key
		reportStep(data, delay, report)
	}
}

func visualQuickSort(data []int, delay time.Duration, report stepReporter) {
	var qs func(low, high int)
	var partition func(low, high int) int

	swap := func(i, j int) {
		if i == j {
			return
		}
		data[i], data[j] = data[j], data[i]
		reportStep(data, delay, report)
	}

	partition = func(low, high int) int {
		pivot := data[high]
		i := low
		for j := low; j < high; j++ {
			if data[j] <= pivot {
				swap(i, j)
				i++
			}
		}
		swap(i, high)
		return i
	}

	qs = func(low, high int) {
		if low < high {
			p := partition(low, high)
			qs(low, p-1)
			qs(p+1, high)
		}
	}

	qs(0, len(data)-1)
}

func visualMergeSort(data []int, delay time.Duration, report stepReporter) {
	if len(data) <= 1 {
		reportStep(data, delay, report)
		return
	}

	buffer := make([]int, len(data))

	var mergeSort func(left, right int)
	var merge func(left, mid, right int)

	merge = func(left, mid, right int) {
		i, j := left, mid
		for k := left; k < right; k++ {
			buffer[k] = data[k]
		}
		for k := left; k < right; k++ {
			if i >= mid {
				data[k] = buffer[j]
				j++
			} else if j >= right {
				data[k] = buffer[i]
				i++
			} else if buffer[i] <= buffer[j] {
				data[k] = buffer[i]
				i++
			} else {
				data[k] = buffer[j]
				j++
			}
			reportStep(data, delay, report)
		}
	}

	mergeSort = func(left, right int) {
		if right-left <= 1 {
			return
		}
		mid := left + (right-left)/2
		mergeSort(left, mid)
		mergeSort(mid, right)
		merge(left, mid, right)
	}

	mergeSort(0, len(data))
}

func visualHeapSort(data []int, delay time.Duration, report stepReporter) {
	n := len(data)
	swap := func(i, j int) {
		data[i], data[j] = data[j], data[i]
		reportStep(data, delay, report)
	}

	var heapify func(n, i int)
	heapify = func(n, i int) {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && data[left] > data[largest] {
			largest = left
		}
		if right < n && data[right] > data[largest] {
			largest = right
		}
		if largest != i {
			swap(i, largest)
			heapify(n, largest)
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(n, i)
	}

	for i := n - 1; i > 0; i-- {
		swap(0, i)
		heapify(i, 0)
	}
}

func visualTimSort(data []int, delay time.Duration, report stepReporter) {
	n := len(data)
	if n <= 1 {
		reportStep(data, delay, report)
		return
	}

	temp := make([]int, n)

	for start := 0; start < n; start += visualMinRun {
		end := start + visualMinRun
		if end > n {
			end = n
		}
		for i := start + 1; i < end; i++ {
			key := data[i]
			j := i - 1
			for j >= start && data[j] > key {
				data[j+1] = data[j]
				j--
				reportStep(data, delay, report)
			}
			data[j+1] = key
			reportStep(data, delay, report)
		}
	}

	for size := visualMinRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size
			right := left + 2*size
			if mid > n {
				mid = n
			}
			if right > n {
				right = n
			}
			if mid >= right {
				continue
			}
			mergeRange(data, temp, left, mid, right, delay, report)
		}
	}
}

func mergeRange(data, temp []int, left, mid, right int, delay time.Duration, report stepReporter) {
	copy(temp[left:right], data[left:right])
	i, j := left, mid
	for k := left; k < right; k++ {
		if i >= mid {
			data[k] = temp[j]
			j++
		} else if j >= right {
			data[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] {
			data[k] = temp[i]
			i++
		} else {
			data[k] = temp[j]
			j++
		}
		reportStep(data, delay, report)
	}
}

func visualRadixSort(data []int, delay time.Duration, report stepReporter) {
	if len(data) <= 1 {
		reportStep(data, delay, report)
		return
	}

	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}

	buffer := make([]int, len(data))
	exp := 1
	for maxVal/exp > 0 {
		count := make([]int, 10)
		for _, v := range data {
			digit := (v / exp) % 10
			count[digit]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for i := len(data) - 1; i >= 0; i-- {
			digit := (data[i] / exp) % 10
			count[digit]--
			buffer[count[digit]] = data[i]
		}
		copy(data, buffer)
		reportStep(data, delay, report)
		exp *= 10
	}
}

func visualCountingSort(data []int, delay time.Duration, report stepReporter) {
	if len(data) == 0 {
		return
	}

	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}

	count := make([]int, maxVal+1)
	for _, v := range data {
		count[v]++
	}

	idx := 0
	for value, freq := range count {
		for freq > 0 {
			data[idx] = value
			idx++
			freq--
			reportStep(data, delay, report)
		}
	}
}

func visualBucketSort(data []int, delay time.Duration, report stepReporter) {
	if len(data) == 0 {
		return
	}

	maxVal := data[0]
	minVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	bucketCount := int(math.Sqrt(float64(len(data))))
	if bucketCount < 1 {
		bucketCount = 1
	}
	buckets := make([][]int, bucketCount)

	rangeVal := maxVal - minVal + 1
	for _, v := range data {
		idx := (v - minVal) * bucketCount / rangeVal
		if idx >= bucketCount {
			idx = bucketCount - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}

	idx := 0
	for _, bucket := range buckets {
		insertionSort(bucket)
		for _, v := range bucket {
			data[idx] = v
			idx++
			reportStep(data, delay, report)
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
