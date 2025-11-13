package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"algosGO/sorts"
)

type algorithm struct {
	Name       string
	Complexity string
	Run        func([]int, sorts.StepCallback)
}

var algorithms = []algorithm{
	{"Быстрая сортировка", "O(n log n)", sorts.QuickSort},
	{"Сортировка слиянием", "O(n log n)", sorts.MergeSort},
	{"Пирамидальная сортировка", "O(n log n)", sorts.HeapSort},
	{"Timsort", "O(n log n)", sorts.TimSort},
	{"Поразрядная сортировка", "O(n)", sorts.RadixSort},
	{"Сортировка подсчётом", "O(n + k)", sorts.CountingSort},
	{"Блочная сортировка", "O(n + k)", sorts.BucketSort},
	{"Сортировка вставками", "O(n^2)", sorts.InsertionSort},
}

const (
	datasetSize = 24
	visualDelay = 120 * time.Millisecond
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выберите алгоритм сортировки:")
	for i, alg := range algorithms {
		fmt.Printf("%d. %s (%s)\n", i+1, alg.Name, alg.Complexity)
	}
	fmt.Println("0. Визуализировать все алгоритмы одновременно")
	fmt.Print("Введите номер и нажмите Enter: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		data := generateData(datasetSize)
		runVisualization(data)
		return
	}

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(algorithms) {
		fmt.Println("Неверный выбор. Перезапустите программу и попробуйте снова.")
		return
	}

	data := generateData(datasetSize)
	runSingle(algorithms[choice-1], data)
}

func generateData(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(900) + 50
	}
	return arr
}

func runSingle(alg algorithm, data []int) {
	arr := append([]int(nil), data...)
	fmt.Printf("\nСтартует %s\n", alg.Name)
	alg.Run(arr, func(state []int) {
		fmt.Println(renderFrame(alg.Name, state))
		time.Sleep(visualDelay)
	})
	fmt.Printf("\n%s завершена. Результат: %v\n", alg.Name, arr)
}

type frame struct {
	name  string
	state []int
	done  bool
}

func runVisualization(data []int) {
	fmt.Println("\nЗапускаем визуализацию всех алгоритмов одновременно...")
	frames := make(chan frame, 64)
	var wg sync.WaitGroup

	for _, alg := range algorithms {
		wg.Add(1)
		go func(a algorithm) {
			defer wg.Done()
			arr := append([]int(nil), data...)
			a.Run(arr, func(state []int) {
				frames <- frame{name: a.Name, state: append([]int(nil), state...)}
				time.Sleep(visualDelay)
			})
			frames <- frame{name: a.Name, state: append([]int(nil), arr...), done: true}
		}(alg)
	}

	go func() {
		wg.Wait()
		close(frames)
	}()

	for fr := range frames {
		if fr.done {
			fmt.Printf("%s завершила сортировку.\n", fr.name)
			continue
		}
		fmt.Println(renderFrame(fr.name, fr.state))
	}
}

func renderFrame(name string, state []int) string {
	if len(state) == 0 {
		return fmt.Sprintf("%s: []", name)
	}
	maxVal := max(state)
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%-26s | ", name))
	for _, v := range state {
		normalized := 1
		if maxVal > 0 {
			normalized = int(math.Round(float64(v) / float64(maxVal) * 8))
			if normalized < 1 {
				normalized = 1
			}
		}
		builder.WriteString(strings.Repeat("▇", normalized))
		builder.WriteByte(' ')
	}
	return builder.String()
}

func max(values []int) int {
	m := values[0]
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}
