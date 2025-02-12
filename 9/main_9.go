package main

import (
	"fmt"
)

// CubePipeline создает конвейер обработки чисел
func CubePipeline(input chan uint8, output chan float64) {
	for num := range input { // Читаем числа из первого канала
		result := float64(num) * float64(num) * float64(num) // Возводим в куб
		output <- result                                     // Записываем результат во второй канал
	}
	close(output) // Закрываем второй канал после завершения работы
}

func main() {
	// Создаем два канала
	input := make(chan uint8)
	output := make(chan float64)

	// Запускаем конвейер в отдельной горутине
	go CubePipeline(input, output)

	// Заполняем первый канал числами
	for i := uint8(1); i <= 5; i++ {
		input <- i
	}
	close(input) // Закрываем первый канал после отправки всех чисел

	// Читаем результаты из второго канала
	fmt.Println("Результаты:")
	for result := range output {
		fmt.Printf("Куб числа: %.2f\n", result)
	}
}
