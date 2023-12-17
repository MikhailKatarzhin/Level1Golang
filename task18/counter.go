package main

import "sync"

// Counter Структура счётчика
type Counter struct {
	mu    sync.Mutex // Мьютекс для защиты доступа к счётчику
	value int        // Значение счётчика
}

// Increment потокобезопасно инкрементирует счётчик
func (c *Counter) Increment() {
	c.mu.Lock()   // Блокируем доступ к счётчику
	c.value++     // Инкрементируем счётчик
	c.mu.Unlock() // Освобождаем счётчик
}
