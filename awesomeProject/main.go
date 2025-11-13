package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"
)

// Структура данных о погоде
type WeatherData struct {
	Date        time.Time `json:"date"`
	Radiation   float64   `json:"radiation"`   // солнечная радиация (Вт/м²)
	Temperature float64   `json:"temperature"` // температура (°C)
	Humidity    float64   `json:"humidity"`    // влажность (%)
	Pressure    float64   `json:"pressure"`    // давление (гПа)
	WindSpeed   float64   `json:"wind_speed"`  // скорость ветра (м/с)
	UVIndex     float64   `json:"uv_index"`    // УФ индекс
}

// Коллекция данных
type WeatherDataset struct {
	Data []WeatherData `json:"data"`
}

// === Методы анализа данных ===

// Среднее значение
func (ds *WeatherDataset) Mean(field string) float64 {
	values := ds.extractField(field)
	return mean(values)
}

// Минимум
func (ds *WeatherDataset) Min(field string) float64 {
	values := ds.extractField(field)
	return min(values)
}

// Максимум
func (ds *WeatherDataset) Max(field string) float64 {
	values := ds.extractField(field)
	return max(values)
}

// Стандартное отклонение
func (ds *WeatherDataset) StdDev(field string) float64 {
	values := ds.extractField(field)
	return stdDev(values)
}

// Корреляция между двумя параметрами
func (ds *WeatherDataset) Correlation(field1, field2 string) float64 {
	x := ds.extractField(field1)
	y := ds.extractField(field2)
	return correlation(x, y)
}

// === Вспомогательные функции ===
func (ds *WeatherDataset) extractField(field string) []float64 {
	values := []float64{}
	for _, d := range ds.Data {
		switch field {
		case "radiation":
			values = append(values, d.Radiation)
		case "temperature":
			values = append(values, d.Temperature)
		case "humidity":
			values = append(values, d.Humidity)
		case "pressure":
			values = append(values, d.Pressure)
		case "wind_speed":
			values = append(values, d.WindSpeed)
		case "uv_index":
			values = append(values, d.UVIndex)
		}
	}
	return values
}

func mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	m := values[0]
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}

func max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	m := values[0]
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func stdDev(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	m := mean(values)
	sum := 0.0
	for _, v := range values {
		sum += (v - m) * (v - m)
	}
	return math.Sqrt(sum / float64(len(values)))
}

func correlation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}
	meanX, meanY := mean(x), mean(y)
	num := 0.0
	denX, denY := 0.0, 0.0

	for i := 0; i < len(x); i++ {
		dx := x[i] - meanX
		dy := y[i] - meanY
		num += dx * dy
		denX += dx * dx
		denY += dy * dy
	}

	return num / math.Sqrt(denX*denY)
}

// === Работа с файлами ===
func (ds *WeatherDataset) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(ds)
}

func LoadFromFile(filename string) (*WeatherDataset, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ds WeatherDataset
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ds)
	return &ds, err
}

// === Ввод данных вручную ===
func inputData() WeatherData {
	var d WeatherData
	d.Date = time.Now()

	fmt.Print("Солнечная радиация (Вт/м²): ")
	fmt.Scan(&d.Radiation)

	fmt.Print("Температура (°C): ")
	fmt.Scan(&d.Temperature)

	fmt.Print("Влажность (%): ")
	fmt.Scan(&d.Humidity)

	fmt.Print("Давление (гПа): ")
	fmt.Scan(&d.Pressure)

	fmt.Print("Скорость ветра (м/с): ")
	fmt.Scan(&d.WindSpeed)

	fmt.Print("УФ индекс: ")
	fmt.Scan(&d.UVIndex)

	return d
}

// === Главная функция ===
func main() {
	const filename = "weather.json"

	// Загружаем данные
	ds, err := LoadFromFile(filename)
	if err != nil {
		ds = &WeatherDataset{}
	}

	// Ввод новых данных
	fmt.Println("Введите новые данные о погоде:")
	newData := inputData()
	ds.Data = append(ds.Data, newData)

	// Сохраняем
	err = ds.SaveToFile(filename)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	// Аналитика
	fmt.Println("\n=== Анализ данных ===")
	fmt.Printf("Средняя температура: %.2f °C\n", ds.Mean("temperature"))
	fmt.Printf("Макс радиация: %.2f\n", ds.Max("radiation"))
	fmt.Printf("Мин влажность: %.2f\n", ds.Min("humidity"))
	fmt.Printf("Стандартное отклонение температуры: %.2f\n", ds.StdDev("temperature"))
	fmt.Printf("Корреляция радиация/температура: %.2f\n", ds.Correlation("radiation", "temperature"))
}
