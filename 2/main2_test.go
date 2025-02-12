package main

import (
	"reflect"
	"testing"
)

// Тест для функции sliceExample
func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

// Тест для функции addElements
func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	valueToAdd := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(input, valueToAdd)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

// Тест для функции copySlice
func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copy := copySlice(input)
	if !reflect.DeepEqual(copy, input) {
		t.Errorf("копия не соответствует оригиналу")
	}

	// Проверка, что изменения в оригинале не влияют на копию
	input[0] = 999
	if reflect.DeepEqual(copy, input) {
		t.Errorf("изменения в оригинале повлияли на копию")
	}
}

// Тест для функции removeElement
func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, indexToRemove)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}
