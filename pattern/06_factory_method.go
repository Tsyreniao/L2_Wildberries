package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Паттерн "фабрика" позволяет легко добавлять новые типы объектов без изменения существующего кода, упрощает управление зависимостями
	и обеспечивает лагическую структуру кода.
	Паттерн может добавить дополнительную сложность при наличии большого кол-ва различных типов продуктов и фабрик.
*/

// Интерфейс продукта
type Product interface {
	Use()
}

// Конкретный продукт 1
type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) Use() {
	fmt.Println("Using ConcreteProduct1")
}

// Конкретный продукт 2
type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) Use() {
	fmt.Println("Using ConcreteProduct2")
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретная фабрика 1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProduct() Product {
	return &ConcreteProduct1{}
}

// Конкретная фабрика 2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProduct() Product {
	return &ConcreteProduct2{}
}

func main() {
	// Создание продукта через фабрику
	factory := &ConcreteFactory1{}
	product := factory.CreateProduct()
	product.Use()

	// Создание другого продукта через другую фабрику
	factory2 := &ConcreteFactory2{}
	product = factory2.CreateProduct()
	product.Use()
}
