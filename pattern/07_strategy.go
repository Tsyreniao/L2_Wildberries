package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Паттерн "стратегия" позволяет разделять ответсвтенность, каждая конкретная стратегия отвечает за свою специфическую реализацию,
	что упрощает код и делает его более поддерживаемым, позволяет легко добавлять новые стратегии.
	Паттерн может добавить дополнительную в организации и управлении стратегиями, при их большом кол-ве.
*/

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64)
}

// Конкретная стратегия: Оплата кредитной картой
type CreditCardPaymentStrategy struct{}

func (s *CreditCardPaymentStrategy) Pay(amount float64) {
	fmt.Printf("Оплата кредитной картой: %.2f\n", amount)
}

// Конкретная стратегия: Оплата через электронный кошелек
type EWalletPaymentStrategy struct{}

func (s *EWalletPaymentStrategy) Pay(amount float64) {
	fmt.Printf("Оплата через электронный кошелек: %.2f\n", amount)
}

// Конкретная стратегия: Оплата наличными
type CashPaymentStrategy struct{}

func (s *CashPaymentStrategy) Pay(amount float64) {
	fmt.Printf("Оплата наличными: %.2f\n", amount)
}

// Контекст, использующий стратегию
type PaymentContext struct {
	strategy PaymentStrategy
}

func (c *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *PaymentContext) MakePayment(amount float64) {
	c.strategy.Pay(amount)
}

func main() {
	// Создаем контекст платежа
	context := &PaymentContext{}

	// Устанавливаем стратегию оплаты кредитной картой
	context.SetPaymentStrategy(&CreditCardPaymentStrategy{})
	context.MakePayment(100.50)

	// Устанавливаем стратегию оплаты через электронный кошелек
	context.SetPaymentStrategy(&EWalletPaymentStrategy{})
	context.MakePayment(50.25)

	// Устанавливаем стратегию оплаты наличными
	context.SetPaymentStrategy(&CashPaymentStrategy{})
	context.MakePayment(75.80)
}
