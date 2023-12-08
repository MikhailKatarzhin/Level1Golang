package task01

import "fmt"

// Структура Human представляет человека
// с одним публичным и одним приватным полем
type Human struct {
	ID   int
	name string
}

// Конструктор для создания экземпляра Human
func NewHuman(newId int, newName string) *Human {
	return &Human{
		ID:   newId,
		name: newName,
	}
}

// Публичный метод IntroduceYourself для структуры Human,
// отправляющий в консоль представление себя
func (h *Human) IntroduceYourself() {
	fmt.Printf("Меня зовут %s, прошу любить и жаловть!", h.giveYourselfName())
}

// Приватный метод giveYourselfName для структуры Human, возвращающий имя строкой
func (h *Human) giveYourselfName() string {
	return h.name
}
