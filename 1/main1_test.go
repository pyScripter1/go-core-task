package main

import (
	"reflect"
	"testing"
)

// Тест для функции конкатенации переменных
func TestConcatenateVariables(t *testing.T) {
	result := concatenateVariables(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	expected := "4242423.14Golangtrue(1+2i)"
	if result != expected {
		t.Errorf("ожидалось '%s', получено '%s'", expected, result)
	}
}

// Тест для функции преобразования строки в руны
func TestStringToRunes(t *testing.T) {
	result := stringToRunes("test")
	expected := []rune{'t', 'e', 's', 't'}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось '%v', получено '%v'", expected, result)
	}
}

// Тест для функции хэширования с солью
func TestHashWithSalt(t *testing.T) {
	runes := stringToRunes("test")
	result := hashWithSalt(runes, "salt")
	expected := "f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0c4feea932c68a56c1c41a2e4"
	if result != expected {
		t.Errorf("ожидалось '%s', получено '%s'", expected, result)
	}
}
