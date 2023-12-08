package task01

import "fmt"

// Action Структура встраивает Human
// и дополняет строкой Всклик
type Action struct {
	Human // встраиваем структуру Human в Action
	Whoop string
}

// NewAction Конструктор для создания экземпляра Action
// на основе имени и идентификатора человека, а также всклика действия
func NewAction(name string, id int, whoop string) *Action {
	return &Action{
		Human: Human{
			ID:   id,
			name: name,
		},
		Whoop: whoop,
	}
}

// DoWhoop Метод для структуры Action пишущий в консоль имя и вскрик
func (a *Action) DoWhoop() {
	fmt.Printf("[Некий %s вскликнул %s!]", a.name, a.Whoop)
}

// GiveYourselfName - метод,
// замещающий одноимённый метод из встроенной структуры Human,
// и возвращающий строкой перевёрнутое Human.name
func (a *Action) GiveYourselfName() string {
	return reverse(a.name)
}

// Метод reverse переворачивает строку
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
