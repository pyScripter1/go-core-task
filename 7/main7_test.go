package main

import (
	"reflect"
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

			// Преобразуем []chan int в []<-chan int
			readOnlyChannels := make([]<-chan int, len(inputChannels))
			for i := range inputChannels {
				readOnlyChannels[i] = inputChannels[i]
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
			merged := MergeChannels(readOnlyChannels...)

			// Читаем данные из объединенного канала
			result := []int{}
			for value := range merged {
				result = append(result, value)
			}

			// Проверяем результат (без учета порядка)
			if !reflect.DeepEqual(sortIntSlice(result), sortIntSlice(tt.expected)) {
				t.Errorf("ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}

// Помощник для сортировки слайса целых чисел
func sortIntSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	for i := 1; i < len(sorted); i++ {
		for j := i; j > 0 && sorted[j-1] > sorted[j]; j-- {
			sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
		}
	}
	return sorted
}
