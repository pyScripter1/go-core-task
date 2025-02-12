package main

import (
	"reflect"
	"testing"
)

// Тест для функции Difference
func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "Простой случай",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "Пустой второй слайс",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "Пустой первый слайс",
			slice1:   []string{},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{},
		},
		{
			name:     "Оба слайса одинаковые",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
		{
			name:     "Пересечение есть, но не все элементы",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"banana"},
			expected: []string{"apple", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}
