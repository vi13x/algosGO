package sorts

import "sort"

// BucketSort реализует блочную сортировку для равномерно распределённых данных.
func BucketSort(arr []int, cb StepCallback) {
	n := len(arr)
	if n == 0 {
		return
	}
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
		emitStep(cb, arr)
		return
	}

	bucketCount := n
	if bucketCount > 50 {
		bucketCount = 50
	}
	buckets := make([][]int, bucketCount)
	rangeVal := maxVal - minVal + 1

	for _, v := range arr {
		idx := (v - minVal) * (bucketCount - 1) / rangeVal
		buckets[idx] = append(buckets[idx], v)
		emitStep(cb, arr)
	}

	pos := 0
	for _, bucket := range buckets {
		sort.Ints(bucket)
		for _, v := range bucket {
			arr[pos] = v
			pos++
			emitStep(cb, arr)
		}
	}
}
