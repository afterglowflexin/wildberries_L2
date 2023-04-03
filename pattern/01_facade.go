// facade - высокоуровневый унифицированный интерфейс
package main

import "fmt"

// facade
type Car struct {
	engine   *Engine
	lights   *Lights
	speakers *Speakers
}

func (c *Car) Start() {
	c.engine.Start()
	c.lights.On()
	c.speakers.On()
}

func NewCar() *Car {
	return &Car{&Engine{}, &Lights{}, &Speakers{}}
}

// subsystem A
type Engine struct {
}

func (e *Engine) Start() {
	fmt.Println("Engine started")
}

// subsystem B
type Lights struct {
}

func (l *Lights) On() {
	fmt.Println("Lights on")
}

// subsystem C
type Speakers struct {
}

func (s *Speakers) On() {
	fmt.Println("Speakers on")
}

func main() {
	car := NewCar()
	car.Start()
}
