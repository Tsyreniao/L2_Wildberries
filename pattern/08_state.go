package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн "состояние" позволяет избежать громоздких условных операторов, позволяет добавлять новое состояние не изменяя существующего кода контекста,
	упрощает понимание и поддержку кода.
	Паттерн может привести к увеличению кол-ва классов, может возникнуть сложность в организации состояний.
*/

// Интерфейс состояния
type State interface {
	Handle(context *Context)
}

// Конкретное состояние: Состояние A
type StateA struct{}

func (s *StateA) Handle(context *Context) {
	fmt.Println("Обработка состояния A")
	// Изменение состояния на следующее
	context.SetState(&StateB{})
}

// Конкретное состояние: Состояние B
type StateB struct{}

func (s *StateB) Handle(context *Context) {
	fmt.Println("Обработка состояния B")
	// Изменение состояния на предыдущее
	context.SetState(&StateA{})
}

// Контекст, использующий состояние
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем начальное состояние
	context.SetState(&StateA{})

	// Выполняем запросы, которые будут обрабатываться в зависимости от текущего состояния
	context.Request()
	context.Request()
	context.Request()
}
