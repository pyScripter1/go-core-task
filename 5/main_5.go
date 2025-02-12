package main

import (
	"fmt"
)

// Intersection проверяет пересечение значений между двумя слайсами и возвращает:
// - bool: есть ли хотя бы одно пересечение
// - []int: срез с пересеченными значениями
func Intersection(slice1, slice2 []int) (bool, []int) {
	set := make(map[int]struct{}) // Множество для хранения уникальных элементов из первого слайса
	var intersection []int        // Срез для хранения пересечений

	// Заполняем множество элементами из первого слайса
	for _, value := range slice1 {
		set[value] = struct{}{}
	}

	// Проверяем второй слайс на наличие элементов из множества
	for _, value := range slice2 {
		if _, exists := set[value]; exists {
			intersection = append(intersection, value)
		}
	}

	// Если найдены пересечения, возвращаем true и срез пересечений
	return len(intersection) > 0, intersection
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, result := Intersection(a, b)
	fmt.Println("Есть пересечение:", hasIntersection)
	fmt.Println("Пересеченные значения:", result)
}
