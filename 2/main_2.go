package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Шаг 1: Создание слайса с 10 случайными числами
	rand.Seed(time.Now().UnixNano())
	originalSlice := generateRandomSlice(10)
	fmt.Println("Исходный слайс:", originalSlice)

	// Шаг 2: Функция sliceExample (фильтрация четных чисел)
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Слайс с четными числами:", evenSlice)

	// Шаг 3: Функция addElements (добавление элемента в конец)
	newSlice := addElements(originalSlice, 42)
	fmt.Println("Слайс после добавления элемента:", newSlice)

	// Шаг 4: Функция copySlice (копирование слайса)
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия исходного слайса:", copiedSlice)

	// Проверка, что изменения в оригинальном слайсе не влияют на копию
	originalSlice[0] = 999
	fmt.Println("Измененный оригинальный слайс:", originalSlice)
	fmt.Println("Копия после изменения оригинала:", copiedSlice)

	// Шаг 5: Функция removeElement (удаление элемента по индексу)
	indexToRemove := 3
	if indexToRemove >= 0 && indexToRemove < len(originalSlice) {
		modifiedSlice := removeElement(originalSlice, indexToRemove)
		fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, modifiedSlice)
	} else {
		fmt.Println("Индекс вне диапазона.")
	}
}

// Генерация слайса с n случайными числами
func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(100) // Числа от 0 до 99
	}
	return slice
}

// Возвращает новый слайс с четными числами из исходного
func sliceExample(slice []int) []int {
	var result []int
	for _, num := range slice {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

// Добавляет число в конец слайса
func addElements(slice []int, value int) []int {
	return append(slice, value)
}

// Создает копию слайса
func copySlice(slice []int) []int {
	copy := make([]int, len(slice))
	copy = append(copy, slice...)
	return copy
}

// Удаляет элемент из слайса по указанному индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Индекс вне диапазона
	}
	return append(slice[:index], slice[index+1:]...)
}
