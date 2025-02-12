package main

import (
	"fmt"
)

// Difference принимает два слайса строк и возвращает новый слайс,
// содержащий элементы из первого слайса, которых нет во втором.
func Difference(slice1, slice2 []string) []string {
	result := []string{}
	set := make(map[string]struct{})

	// Заполняем множество элементами из второго слайса
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	// Проверяем каждый элемент первого слайса
	for _, item := range slice1 {
		if _, exists := set[item]; !exists {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := Difference(slice1, slice2)
	fmt.Println("Разница:", result)
}
