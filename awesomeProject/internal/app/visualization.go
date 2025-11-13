package app

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	visualColumns   = 24
	visualMaxHeight = 400
	visualStepDelay = 75 * time.Millisecond
	visualMinRun    = 32
)

var blockRunes = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

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
