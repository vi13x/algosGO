package sorts

import "sort"

// BucketSort выполняет блочную сортировку и возвращает отсортированную копию слайса.
func BucketSort(input []int) []int {
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

	if minVal == maxVal {
		return arr
	}

	bucketCount := len(arr)
	buckets := make([][]int, bucketCount)
	rangeVal := maxVal - minVal + 1

	for _, v := range arr {
		idx := (v - minVal) * (bucketCount - 1) / rangeVal
		buckets[idx] = append(buckets[idx], v)
	}

	idx := 0
	for _, bucket := range buckets {
		if len(bucket) == 0 {
			continue
		}
		sort.Ints(bucket)
		for _, v := range bucket {
			arr[idx] = v
			idx++
		}
	}

	return arr
}
