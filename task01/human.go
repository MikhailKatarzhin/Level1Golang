package task01

import "fmt"

// Human Структура представляет человека
// с одним "публичным" и одним "приватным" полем
type Human struct {
	ID   int
	name string
}

// NewHuman Конструктор для создания экземпляра
func NewHuman(newId int, newName string) *Human {
	return &Human{
		ID:   newId,
		name: newName,
	}
}

// IntroduceYourself "Публичный" метод  для структуры Human,
// отправляющий в консоль представление себя
func (h *Human) IntroduceYourself() {
	fmt.Printf("Меня зовут %s, прошу любить и жаловть!", h.GiveYourselfName())
}

// GiveYourselfName "Публичный" метод  для структуры Human, возвращающий имя строкой
func (h *Human) GiveYourselfName() string {
	return h.name
}
