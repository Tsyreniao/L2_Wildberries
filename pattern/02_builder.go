package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Паттерн "строитель" упрощает содание сложных объектов путем пошагового конструирования, улучшает читаемость и понимание кода,
	разделение процесса создания объекта делает конструкцию более гибкой.
	Паттерн может усложнить код, конструирование объекта не является сложным и не требует множества шагов.
*/

// Product - конечный продукт, объект который будет создан с помощью паттерна "Строитель"
type Product struct {
	partA string
	partB string
	partC string
}

// Builder - интерфейс для строителя, определяющий шаги конструирования
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// ConcreteBuilder - конкретный строитель, реализующий интерфейс Builder
type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "Part B"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "Part C"
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Director - директор, который управляет процессом конструирования
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}

	director.Construct()
	product := builder.GetProduct()

	fmt.Println(product)
}
