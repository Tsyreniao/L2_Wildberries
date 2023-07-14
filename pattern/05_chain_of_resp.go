package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Паттерн "цепочка вызовов" разделяет обязанности, каждый объект в цепочке отвечает только за свою часть, упрощает поддержку и расширение системы,
	объекты в цепочке не зависят друг от друга напрямую, что уменьшает связаность системы.
	Паттерн не гарантирует обработку запроса, при наличии большого числа обработчиков в цепи может возникнуть задержка.
*/

// Объявление интерфейса для обработчика запросов
type Handler interface {
	SetNext(handler Handler)      // Установка следующего обработчика в цепочке
	HandleRequest(request string) // Обработка запроса
}

// Конкретный обработчик запросов
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA обрабатывает запрос", request)
	} else if h.next != nil {
		h.next.HandleRequest(request) // Передача запроса следующему обработчику
	}
}

// Конкретный обработчик запросов
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB обрабатывает запрос", request)
	} else if h.next != nil {
		h.next.HandleRequest(request) // Передача запроса следующему обработчику
	}
}

// Конкретный обработчик запросов
type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerC) HandleRequest(request string) {
	if request == "C" {
		fmt.Println("ConcreteHandlerC обрабатывает запрос", request)
	} else if h.next != nil {
		h.next.HandleRequest(request) // Передача запроса следующему обработчику
	}
}

func main() {
	// Создание цепочки обработчиков
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}
	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	// Обработка запросов
	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
	handlerA.HandleRequest("D")
}
