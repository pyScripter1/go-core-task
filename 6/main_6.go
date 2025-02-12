package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumberGenerator - функция-генератор случайных чисел
func RandomNumberGenerator(min, max int) <-chan int {
	ch := make(chan int) // Создаем небуферизированный канал
	go func() {
		defer close(ch)                  // Закрываем канал после завершения работы горутины
		rand.Seed(time.Now().UnixNano()) // Инициализируем рандомизатор
		for {
			ch <- rand.Intn(max-min+1) + min // Отправляем случайное число в канал
		}
	}()
	return ch
}

// GetRandomNumbers - получает указанное количество случайных чисел из генератора
func GetRandomNumbers(generator <-chan int, count int) []int {
	var numbers []int
	for i := 0; i < count; i++ {
		num := <-generator // Получаем число из канала
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	// Создаем генератор случайных чисел от 1 до 100
	generator := RandomNumberGenerator(1, 100)

	// Получаем 10 случайных чисел
	randomNumbers := GetRandomNumbers(generator, 10)

	// Выводим результат
	fmt.Println("Случайные числа:", randomNumbers)
}
