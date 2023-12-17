package main

// LegacyObject представляющий объект, с которым мы хотим работать через адаптер
type LegacyObject struct{}

func (o *LegacyObject) specificRequest() string {
	// неизвестные обработки старого объекта до отправки ответа
	return "LegacyObject: specificRequest"
}
