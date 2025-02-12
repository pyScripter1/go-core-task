package main

import (
	"reflect"
	"testing"
)

// Тест для функции Intersection
func TestIntersection(t *testing.T) {
	tests := []struct {
		name          string
		slice1        []int
		slice2        []int
		expectedBool  bool
		expectedSlice []int
	}{
		{
			name:          "Простой случай",
			slice1:        []int{65, 3, 58, 678, 64},
			slice2:        []int{64, 2, 3, 43},
			expectedBool:  true,
			expectedSlice: []int{64, 3},
		},
		{
			name:          "Нет пересечений",
			slice1:        []int{1, 2, 3},
			slice2:        []int{4, 5, 6},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Пустой первый слайс",
			slice1:        []int{},
			slice2:        []int{4, 5, 6},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Пустой второй слайс",
			slice1:        []int{1, 2, 3},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Оба слайса пустые",
			slice1:        []int{},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Все элементы совпадают",
			slice1:        []int{1, 2, 3},
			slice2:        []int{1, 2, 3},
			expectedBool:  true,
			expectedSlice: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultBool, resultSlice := Intersection(tt.slice1, tt.slice2)

			if resultBool != tt.expectedBool {
				t.Errorf("ожидалось %v, получено %v", tt.expectedBool, resultBool)
			}

			if !reflect.DeepEqual(resultSlice, tt.expectedSlice) {
				t.Errorf("ожидалось %v, получено %v", tt.expectedSlice, resultSlice)
			}
		})
	}
}
