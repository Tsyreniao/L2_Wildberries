package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Паттерн "посетитель" упрощает добавление нового вункционала, улучшает читаемость и поддерживаемость кода,
	позволяет работать с различными типами объектов через общий интерфейс посетителя.
	Паттерн может увеличить сложность кода, про работе с большым кол-вом различных классов, паттерн может нарушить инкапсуляцию.
*/

// Интерфейс элемента, принимающего посетителя
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент
type ConcreteElementA struct{}

func (ce *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(ce)
}

// Конкретный элемент
type ConcreteElementB struct{}

func (ce *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ce)
}

// Интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("VisitConcreteElementA visited")
}

func (cv *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("VisitConcreteElementB visited")
}

// Клиентский код
func main() {
	visitor := &ConcreteVisitor{}

	elementA := &ConcreteElementA{}
	elementA.Accept(visitor)

	elementB := &ConcreteElementB{}
	elementB.Accept(visitor)
}
