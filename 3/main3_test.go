package main

import (
	"reflect"
	"testing"
)

// Тест для метода Add
func TestAdd(t *testing.T) {
	mapInstance := NewStringIntMap()
	mapInstance.Add("test", 42)

	expected := map[string]int{"test": 42}
	if !reflect.DeepEqual(mapInstance.data, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, mapInstance.data)
	}
}

// Тест для метода Remove
func TestRemove(t *testing.T) {
	mapInstance := NewStringIntMap()
	mapInstance.Add("test", 42)
	mapInstance.Remove("test")

	expected := map[string]int{}
	if !reflect.DeepEqual(mapInstance.data, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, mapInstance.data)
	}
}

// Тест для метода Copy
func TestCopy(t *testing.T) {
	mapInstance := NewStringIntMap()
	mapInstance.Add("test", 42)

	copied := mapInstance.Copy()
	expected := map[string]int{"test": 42}

	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, copied)
	}

	// Изменение оригинальной карты не должно влиять на копию
	mapInstance.Add("newKey", 99)
	if reflect.DeepEqual(copied, mapInstance.data) {
		t.Errorf("копия должна быть независимой от оригинала")
	}
}

// Тест для метода Exists
func TestExists(t *testing.T) {
	mapInstance := NewStringIntMap()
	mapInstance.Add("test", 42)

	if !mapInstance.Exists("test") {
		t.Errorf("ключ 'test' должен существовать")
	}

	if mapInstance.Exists("nonexistent") {
		t.Errorf("ключ 'nonexistent' не должен существовать")
	}
}

// Тест для метода Get
func TestGet(t *testing.T) {
	mapInstance := NewStringIntMap()
	mapInstance.Add("test", 42)

	if value, ok := mapInstance.Get("test"); !ok || value != 42 {
		t.Errorf("ожидалось значение 42, получено %d", value)
	}

	if _, ok := mapInstance.Get("nonexistent"); ok {
		t.Errorf("ключ 'nonexistent' не должен существовать")
	}
}
