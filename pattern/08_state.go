// Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния
// и является объектно-ориентированной реализацией конечного автомата.
//
// поведение объекта зависит от его состояния
// поведение объекта должно изменяться во время выполнения программы
// состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно

package main

// интерфейс стейтера
type EngineModeStater interface {
	Drive() string
}

// объект с состоянием
type Car struct {
	state EngineModeStater
}

// метод установки состояния
func (c *Car) SetMode(state EngineModeStater) {
	c.state = state
}

// метод работы исходя из состояния
func (c *Car) Drive() string {
	return c.state.Drive()
}

// конструктор объекта
func NewCar() *Car {
	return &Car{state: ComfortMode{}}
}

// конкретный стейт/состояние №1
type ComfortMode struct {
}

func (c ComfortMode) Drive() string {
	return "Driving at 80 km/h"
}

// конкретный стейт/состояние №2
type SportMode struct {
}

func (c SportMode) Drive() string {
	return "Driving at 160 km/h"
}

// конкретный стейт/состояние №3
type EcoMode struct {
}

func (c EcoMode) Drive() string {
	return "Driving at 40 km/h"
}
