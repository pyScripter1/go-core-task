package main

import (
	"testing"
)

// Тест для CubePipeline
func TestCubePipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []float64
	}{
		{
			name:     "Простой случай",
			input:    []uint8{1, 2, 3},
			expected: []float64{1.0, 8.0, 27.0},
		},
		{
			name:     "Большое количество чисел",
			input:    []uint8{4, 5, 6, 7},
			expected: []float64{64.0, 125.0, 216.0, 343.0},
		},
		{
			name:     "Пустой вход",
			input:    []uint8{},
			expected: []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем каналы
			input := make(chan uint8)
			output := make(chan float64)

			// Запускаем конвейер в отдельной горутине
			go CubePipeline(input, output)

			// Отправляем данные в первый канал
			for _, num := range tt.input {
				input <- num
			}
			close(input)

			// Считываем результаты из второго канала
			var results []float64
			for result := range output {
				results = append(results, result)
			}

			// Проверяем результаты
			if len(results) != len(tt.expected) {
				t.Errorf("ожидалось %d элементов, получено %d", len(tt.expected), len(results))
			}
			for i, v := range results {
				if v != tt.expected[i] {
					t.Errorf("ожидалось %.2f, получено %.2f", tt.expected[i], v)
				}
			}
		})
	}
}
