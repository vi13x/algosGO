package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Run стартует приложение, предлагая пользователю выбрать режим работы.
func Run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Выберите режим:")
	fmt.Println("1 - Отсортировать введённые числа выбранным алгоритмом")
	fmt.Println("2 - Визуализировать работу всех алгоритмов одновременно")
	fmt.Print("Ваш выбор: ")

	modeLine, err := readLine(reader)
	if err != nil && err != io.EOF {
		return fmt.Errorf("не удалось прочитать выбор: %w", err)
	}
	if modeLine == "" {
		return fmt.Errorf("не указан режим работы")
	}

	mode, err := strconv.Atoi(modeLine)
	if err != nil {
		return fmt.Errorf("ожидалось число, которое задаёт режим")
	}

	switch mode {
	case 1:
		runInteractive(reader)
	case 2:
		runVisualization()
	default:
		return fmt.Errorf("неизвестный режим, завершение работы")
	}

	return nil
}
