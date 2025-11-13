package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"awesomeProject/sorts"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите целые числа через пробел:")
	line, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Не удалось прочитать ввод:", err)
		return
	}

	fields := strings.Fields(strings.TrimSpace(line))
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

	var choice int
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Ошибка чтения выбора:", err)
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
