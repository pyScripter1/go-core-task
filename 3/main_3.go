package main

import (
	"fmt"
)

// StringIntMap представляет собой структуру данных для хранения пар "строка - целое число"
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap создает новую карту
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет новую пару "ключ-значение" в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу из карты
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает копию текущей карты
func (m *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int)
	for k, v := range m.data {
		copied[k] = v
	}
	return copied
}

// Exists проверяет, существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Get возвращает значение по ключу и булевый флаг успешности операции
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

// Метод для вывода содержимого карты (для отладки)
func (m *StringIntMap) Print() {
	fmt.Println("Содержимое карты:")
	for k, v := range m.data {
		fmt.Printf("Ключ: %s, Значение: %d\n", k, v)
	}
}

func main() {
	// Создание новой карты
	mapInstance := NewStringIntMap()

	// Добавление элементов
	mapInstance.Add("apple", 10)
	mapInstance.Add("banana", 20)
	mapInstance.Add("cherry", 30)

	// Вывод содержимого карты
	mapInstance.Print()

	// Проверка наличия ключа
	fmt.Println("Exists('banana'):", mapInstance.Exists("banana"))
	fmt.Println("Exists('grape'):", mapInstance.Exists("grape"))

	// Получение значения
	if value, ok := mapInstance.Get("apple"); ok {
		fmt.Println("Get('apple'):", value)
	} else {
		fmt.Println("Ключ 'apple' не найден")
	}

	// Копирование карты
	copiedMap := mapInstance.Copy()
	fmt.Println("Содержимое скопированной карты:")
	for k, v := range copiedMap {
		fmt.Printf("Ключ: %s, Значение: %d\n", k, v)
	}

	// Удаление элемента
	mapInstance.Remove("banana")
	fmt.Println("После удаления 'banana':")
	mapInstance.Print()
}
