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

	fmt.Println("Выберите алгоритм сортировки:")
	fmt.Println("1 - Сортировка вставками")
	fmt.Println("2 - Быстрая сортировка")
	fmt.Print("Ваш выбор: ")

	var choice int
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Ошибка чтения выбора:", err)
		return
	}

	var result []int
	switch choice {
	case 1:
		result = sorts.InsertionSort(numbers)
		fmt.Println("Вы выбрали сортировку вставками.")
	case 2:
		result = sorts.QuickSort(numbers)
		fmt.Println("Вы выбрали быструю сортировку.")
	default:
		fmt.Println("Неизвестный выбор, завершение работы.")
		return
	}

	fmt.Printf("Исходный массив: %v\n", numbers)
	fmt.Printf("Отсортированный массив: %v\n", result)
}
