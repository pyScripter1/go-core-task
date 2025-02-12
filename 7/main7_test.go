package main

import (
	"testing"
)

// Тест для MergeChannels
func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name     string
		channels [][]int
		expected []int
	}{
		{
			name: "Три канала",
			channels: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, // Порядок может отличаться
		},
		{
			name: "Один канал",
			channels: [][]int{
				{10, 20, 30},
			},
			expected: []int{10, 20, 30},
		},
		{
			name: "Пустые каналы",
			channels: [][]int{
				{},
				{},
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем каналы для теста
			inputChannels := make([]chan int, len(tt.channels))
			for i := range inputChannels {
				inputChannels[i] = make(chan int)
			}

			// Запускаем горутины для отправки данных в каналы
			for i, values := range tt.channels {
				go func(ch chan int, vals []int) {
					defer close(ch)
					for _, v := range vals {
						ch <- v
					}
				}(inputChannels[i], values)
			}

			// Сливаем каналы
			merged := MergeChannels(inputChannels...)

			// Читаем данные из объединенного канала
			result := []int{}
			for value := range merged {
				result = append(result, value)
			}

			// Проверяем результат (без учета порядка)
			if len(result) != len(tt.expected) {
				t.Errorf("ожидалось %d элементов, получено %d", len(tt.expected), len(result))
			}
			for _, v := range result {
				if !contains(tt.expected, v) {
					t.Errorf("элемент %d не найден в ожидаемых значениях", v)
				}
			}
		})
	}
}

// Помощник для проверки наличия элемента в слайсе
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
