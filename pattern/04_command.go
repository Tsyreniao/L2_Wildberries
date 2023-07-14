package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Паттерн "команда" разделяет отправителя команды от получателя, что обеспечивает более слабую связанность м/у объектами,
	облегчает расширение системы.
	Паттерн может увеличить сложность кода из-за необходимости создания дополнительных классов, необходимость аккуратной организации иерархии команд.
*/

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is off")
}

// Concrete Command
type TurnOnCommand struct {
	Light *Light
}

func (c *TurnOnCommand) Execute() {
	c.Light.TurnOn()
}

// Concrete Command
type TurnOffCommand struct {
	Light *Light
}

func (c *TurnOffCommand) Execute() {
	c.Light.TurnOff()
}

// Invoker
type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

func main() {
	// Create Receiver
	light := &Light{}

	// Create Concrete Commands
	turnOnCommand := &TurnOnCommand{Light: light}
	turnOffCommand := &TurnOffCommand{Light: light}

	// Create Invoker
	remoteControl := &RemoteControl{}

	// Set Concrete Command to Invoker
	remoteControl.Command = turnOnCommand
	remoteControl.PressButton()

	remoteControl.Command = turnOffCommand
	remoteControl.PressButton()
}
