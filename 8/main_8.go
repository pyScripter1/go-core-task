package main

import (
	"fmt"
)

// CustomWaitGroup представляет собой кастомную реализацию WaitGroup
type CustomWaitGroup struct {
	semaphore chan struct{} // Семафор для синхронизации
}

// NewCustomWaitGroup создает новый экземпляр CustomWaitGroup
func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}),
	}
}

// Add увеличивает счетчик задач на n
func (cw *CustomWaitGroup) Add(n int) {
	for i := 0; i < n; i++ {
		cw.semaphore <- struct{}{}
	}
}

// Done уменьшает счетчик задач на 1
func (cw *CustomWaitGroup) Done() {
	<-cw.semaphore
}

// Wait блокирует выполнение до тех пор, пока счетчик задач не станет равным нулю
func (cw *CustomWaitGroup) Wait() {
	for len(cw.semaphore) > 0 {
		<-cw.semaphore
	}
}

func main() {
	// Создаем новый экземпляр CustomWaitGroup
	cwg := NewCustomWaitGroup()

	// Запускаем горутины
	for i := 1; i <= 5; i++ {
		cwg.Add(1) // Увеличиваем счетчик задач
		go func(id int) {
			defer cwg.Done() // Уменьшаем счетчик задач после завершения работы
			fmt.Printf("Горутина %d работает\n", id)
		}(i)
	}

	// Ожидаем завершения всех горутин
	fmt.Println("Ожидание завершения всех горутин...")
	cwg.Wait()
	fmt.Println("Все горутины завершили работу.")
}
