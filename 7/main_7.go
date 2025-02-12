package main

import (
	"fmt"
)

// MergeChannels сливает N каналов в один выходной канал.
// Возвращает закрытый канал после того, как все входные каналы будут закрыты.
func MergeChannels(channels ...<-chan int) <-chan int {
	mergedChannel := make(chan int)

	go func() {
		defer close(mergedChannel) // Закрываем выходной канал после завершения работы

		// Создаем выборку для всех входных каналов
		var activeChannels []<-chan int = channels

		for len(activeChannels) > 0 {
			select {
			case value, ok := <-activeChannels[0]:
				if ok {
					mergedChannel <- value // Пересылаем значение в выходной канал
				} else {
					// Если канал закрыт, удаляем его из списка активных каналов
					activeChannels = append(activeChannels[:0], activeChannels[1:]...)
				}
			default:
				// Если первый канал не готов к чтению, переключаемся на следующий
				newActiveChannels := activeChannels[1:]
				if len(newActiveChannels) > 0 {
					activeChannels = append(newActiveChannels, activeChannels[0])
				} else {
					break
				}
			}
		}
	}()

	return mergedChannel
}

func main() {
	// Создаем три канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 7; i <= 9; i++ {
			ch3 <- i
		}
	}()

	// Сливаем каналы
	merged := MergeChannels(ch1, ch2, ch3)

	// Читаем данные из объединенного канала
	fmt.Println("Слияние каналов:")
	for value := range merged {
		fmt.Println(value)
	}
}
