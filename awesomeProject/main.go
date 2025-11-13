package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"awesomeProject/sorts"
)

const (
	visualColumns   = 24
	visualMaxHeight = 400
	visualStepDelay = 75 * time.Millisecond
	visualMinRun    = 32
)

var blockRunes = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Выберите режим:")
	fmt.Println("1 - Отсортировать введённые числа выбранным алгоритмом")
	fmt.Println("2 - Визуализировать работу всех алгоритмов одновременно")
	fmt.Print("Ваш выбор: ")

	modeLine, err := readLine(reader)
	if err != nil && err != io.EOF {
		fmt.Println("Не удалось прочитать выбор:", err)
		return
	}
	if modeLine == "" {
		fmt.Println("Не указан режим работы")
		return
	}
	mode, err := strconv.Atoi(modeLine)
	if err != nil {
		fmt.Println("Ожидалось число, которое задаёт режим")
		return
	}

	switch mode {
	case 1:
		runInteractive(reader)
	case 2:
		runVisualization()
	default:
		fmt.Println("Неизвестный режим, завершение работы")
	}
}

func runInteractive(reader *bufio.Reader) {
	fmt.Println("Введите целые числа через пробел:")
	line, err := readLine(reader)
	if err != nil && err != io.EOF {
		fmt.Println("Не удалось прочитать числа:", err)
		return
	}

	fields := strings.Fields(line)
	if len(fields) == 0 {
		fmt.Println("Список чисел пуст")
		return
	}

	numbers := make([]int, 0, len(fields))
	for _, f := range fields {
		value, err := strconv.Atoi(f)
		if err != nil {
			fmt.Printf("%q не является целым числом\n", f)
			return
		}
		numbers = append(numbers, value)
	}

	algorithms := []struct {
		name   string
		sorter func([]int) []int
	}{
		{"Сортировка вставками", sorts.InsertionSort},
		{"Быстрая сортировка", sorts.QuickSort},
		{"Сортировка слиянием", sorts.MergeSort},
		{"Пирамидальная сортировка", sorts.HeapSort},
		{"Timsort", sorts.TimSort},
		{"Поразрядная сортировка", sorts.RadixSort},
		{"Сортировка подсчётом", sorts.CountingSort},
		{"Блочная сортировка", sorts.BucketSort},
	}

	fmt.Println("Выберите алгоритм сортировки:")
	for i, alg := range algorithms {
		fmt.Printf("%d - %s\n", i+1, alg.name)
	}
	fmt.Print("Ваш выбор: ")

	choiceLine, err := readLine(reader)
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения выбора:", err)
		return
	}
	choice, err := strconv.Atoi(strings.TrimSpace(choiceLine))
	if err != nil {
		fmt.Println("Ожидалось число для выбора алгоритма")
		return
	}

	if choice < 1 || choice > len(algorithms) {
		fmt.Println("Неизвестный выбор, завершение работы.")
		return
	}

	selected := algorithms[choice-1]
	result := selected.sorter(numbers)

	fmt.Printf("Вы выбрали: %s\n", selected.name)
	fmt.Printf("Исходный массив: %v\n", numbers)
	fmt.Printf("Отсортированный массив: %v\n", result)
}

func runVisualization() {
	fmt.Println("Запускаю визуализацию. Используются автоматически сгенерированные столбцы.")
	data := generateColumns(visualColumns, visualMaxHeight)
	fmt.Printf("Начальные значения: %v\n", data)
	fmt.Println("Каждый алгоритм работает со своей копией данных. Нажмите Ctrl+C, чтобы прервать процесс.")

	updates := make(chan visualUpdate)
	var wg sync.WaitGroup
	algorithms := visualAlgorithms()

	for _, alg := range algorithms {
		wg.Add(1)
		go func(alg visualAlgorithm) {
			defer wg.Done()
			arr := append([]int(nil), data...)
			sendVisualUpdate(updates, alg.name, arr, false)
			alg.run(arr, visualStepDelay, func(state []int) {
				sendVisualUpdate(updates, alg.name, state, false)
			})
			sendVisualUpdate(updates, alg.name, arr, true)
		}(alg)
	}

	go func() {
		wg.Wait()
		close(updates)
	}()

	renderVisualUpdates(updates, algorithms, visualMaxHeight)
}

type visualAlgorithm struct {
	name string
	run  func([]int, time.Duration, stepReporter)
}

type visualUpdate struct {
	name string
	data []int
	done bool
}

type stepReporter func([]int)

func visualAlgorithms() []visualAlgorithm {
	return []visualAlgorithm{
		{"Сортировка вставками", visualInsertionSort},
		{"Быстрая сортировка", visualQuickSort},
		{"Сортировка слиянием", visualMergeSort},
		{"Пирамидальная сортировка", visualHeapSort},
		{"Timsort", visualTimSort},
		{"Поразрядная сортировка", visualRadixSort},
		{"Сортировка подсчётом", visualCountingSort},
		{"Блочная сортировка", visualBucketSort},
	}
}

func sendVisualUpdate(ch chan<- visualUpdate, name string, data []int, done bool) {
	snapshot := append([]int(nil), data...)
	ch <- visualUpdate{name: name, data: snapshot, done: done}
}

func renderVisualUpdates(updates <-chan visualUpdate, algorithms []visualAlgorithm, maxValue int) {
	states := make(map[string][]int)
	finished := make(map[string]bool)

	for update := range updates {
		states[update.name] = update.data
		if update.done {
			finished[update.name] = true
		}
		clearScreen()
		fmt.Println("Визуализация сортировок (▁ -> низкий столбец, █ -> высокий):")
		fmt.Println()
		for _, alg := range algorithms {
			status := "в процессе"
			if finished[alg.name] {
				status = "готово"
			}
			fmt.Printf("%s [%s]\n", alg.name, status)
			if state, ok := states[alg.name]; ok {
				fmt.Println(renderColumns(state, maxValue))
			} else {
				fmt.Println("(ожидание данных)")
			}
			fmt.Println()
		}
	}

	fmt.Println("Визуализация завершена.")
}

func renderColumns(data []int, maxValue int) string {
	if maxValue == 0 || len(data) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, v := range data {
		idx := 0
		if v > 0 {
			idx = int(math.Round(float64(v) / float64(maxValue) * float64(len(blockRunes)-1)))
			if idx >= len(blockRunes) {
				idx = len(blockRunes) - 1
			}
		}
		builder.WriteRune(blockRunes[idx])
		builder.WriteRune(' ')
	}
	return builder.String()
}

func generateColumns(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Intn(max-50) + 50
	}
	return result
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(line), err
}

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
