package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"awesomeProject/sorts"
)

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

	algorithms := availableAlgorithms()

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

type algorithm struct {
	name   string
	sorter func([]int) []int
}

func availableAlgorithms() []algorithm {
	return []algorithm{
		{"Сортировка вставками", sorts.InsertionSort},
		{"Быстрая сортировка", sorts.QuickSort},
		{"Сортировка слиянием", sorts.MergeSort},
		{"Пирамидальная сортировка", sorts.HeapSort},
		{"Timsort", sorts.TimSort},
		{"Поразрядная сортировка", sorts.RadixSort},
		{"Сортировка подсчётом", sorts.CountingSort},
		{"Блочная сортировка", sorts.BucketSort},
	}
}
