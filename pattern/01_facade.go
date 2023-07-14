package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Паттерн "фасад" упрощает взаимодействие с подсистемой и скрывает его сложность, повышает уровень абстракции кода,
	улучшает читаемость.
	Паттерн может привести к появлению большого и неразборчивого фасада, также сокрытие сложности может привести к потере гибкости.
*/

// Сложная подсистема

type Subsystem1 struct{}

func (s *Subsystem1) Operation1() {
	fmt.Println("Subsystem1: Operation1")
}

type Subsystem2 struct{}

func (s *Subsystem2) Operation2() {
	fmt.Println("Subsystem2: Operation2")
}

type Subsystem3 struct{}

func (s *Subsystem3) Operation3() {
	fmt.Println("Subsystem3: Operation3")
}

// Фасад

type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
	subsystem3 *Subsystem3
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
		subsystem3: &Subsystem3{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.subsystem1.Operation1()
	f.subsystem2.Operation2()
	f.subsystem3.Operation3()
}

// Клиентский код

func main() {
	facade := NewFacade()
	facade.Operation()
}
