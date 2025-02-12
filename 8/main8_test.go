package main

import (
	"sync"
	"testing"
	"time"
)

// Тест для CustomWaitGroup
func TestCustomWaitGroup(t *testing.T) {
	tests := []struct {
		name     string
		tasks    int
		expected int
	}{
		{
			name:     "Проверка базового случая",
			tasks:    5,
			expected: 0,
		},
		{
			name:     "Большое количество задач",
			tasks:    100,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем новый экземпляр CustomWaitGroup
			cwg := NewCustomWaitGroup()

			// Синхронизация для ожидания завершения тестовых горутин
			var wg sync.WaitGroup

			// Запускаем заданное количество задач
			for i := 1; i <= tt.tasks; i++ {
				wg.Add(1)
				cwg.Add(1)
				go func(id int) {
					defer wg.Done()
					defer cwg.Done()
					time.Sleep(10 * time.Millisecond) // Имитация работы
				}(i)
			}

			// Запускаем таймер для проверки времени ожидания
			done := make(chan struct{})
			go func() {
				cwg.Wait()
				close(done)
			}()

			// Ждем завершения всех задач
			wg.Wait()

			// Проверяем, что все задачи завершились
			select {
			case <-done:
				// Все хорошо, канал закрыт
			case <-time.After(1 * time.Second):
				t.Errorf("ожидалось завершение в течение 1 секунды")
			}
		})
	}
}
