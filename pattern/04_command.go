// команда - это объект. Receiver - объект-получатель запроса. Invoker - объект-инициатор запроса
package main

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

type Receiver struct {
}

func (r *Receiver) ToggleOn() string {
	return "Toggle on"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle off"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) DeleteCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) ExecuteAll() {
	for _, c := range i.commands {
		c.Execute()
	}
}
