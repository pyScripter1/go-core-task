package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	// Шаг 1: Создание переменных различных типов
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	// Шаг 2: Определение типа каждой переменной и вывод его на экран
	printVariableType("numDecimal", numDecimal)
	printVariableType("numOctal", numOctal)
	printVariableType("numHexadecimal", numHexadecimal)
	printVariableType("pi", pi)
	printVariableType("name", name)
	printVariableType("isActive", isActive)
	printVariableType("complexNum", complexNum)

	// Шаг 3: Преобразование всех переменных в строковый тип и объединение их в одну строку
	str := concatenateVariables(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Объединенная строка:", str)

	// Шаг 4: Преобразование строки в срез рун
	runes := stringToRunes(str)
	fmt.Println("Срез рун:", runes)

	// Шаг 5: Хэширование среза рун SHA256 с добавлением соли
	hashed := hashWithSalt(runes, "go-2024")
	fmt.Println("Хэш SHA256:", hashed)
}

// Функция для вывода типа переменной
func printVariableType(name string, value interface{}) {
	fmt.Printf("%s имеет тип %s\n", name, reflect.TypeOf(value))
}

// Функция для объединения переменных в одну строку
func concatenateVariables(variables ...interface{}) string {
	var result string
	for _, v := range variables {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

// Функция для преобразования строки в срез рун
func stringToRunes(s string) []rune {
	return []rune(s)
}

// Функция для хэширования среза рун с добавлением соли
func hashWithSalt(runes []rune, salt string) string {
	// Преобразование среза рун в строку
	data := string(runes)

	// Добавление соли в середину строки
	mid := len(data) / 2
	hashInput := data[:mid] + salt + data[mid:]

	// Хэширование SHA256
	hash := sha256.Sum256([]byte(hashInput))
	return fmt.Sprintf("%x", hash)
}
