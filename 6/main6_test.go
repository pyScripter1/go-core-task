package main

import (
	"testing"
	"time"
)

// Тест для RandomNumberGenerator и GetRandomNumbers
func TestRandomNumberGenerator(t *testing.T) {
	tests := []struct {
		name     string
		min      int
		max      int
		count    int
		expected bool
	}{
		{
			name:     "Проверка диапазона",
			min:      1,
			max:      100,
			count:    10,
			expected: true,
		},
		{
			name:     "Маленький диапазон",
			min:      5,
			max:      10,
			count:    5,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем генератор
			generator := RandomNumberGenerator(tt.min, tt.max)

			// Получаем случайные числа
			numbers := GetRandomNumbers(generator, tt.count)

			// Проверяем, что все числа находятся в диапазоне [min, max]
			for _, num := range numbers {
				if num < tt.min || num > tt.max {
					t.Errorf("число %d выходит за пределы диапазона [%d, %d]", num, tt.min, tt.max)
				}
			}

			// Проверяем, что количество полученных чисел равно заданному
			if len(numbers) != tt.count {
				t.Errorf("ожидалось %d чисел, получено %d", tt.count, len(numbers))
			}
		})
	}

	// Тест на закрытие канала (генератор должен работать бесконечно)
	t.Run("Закрытие канала", func(t *testing.T) {
		generator := RandomNumberGenerator(1, 100)
		select {
		case <-generator:
			// Канал открыт, тест успешен
		case <-time.After(1 * time.Second):
			t.Error("канал закрылся раньше времени")
		}
	})
}
