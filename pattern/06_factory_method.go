// Паттерн Factory Method применяется для создания объектов с определенным интерфейсом,
// реализации которого предоставляются потомками.
//
// Factory Method отличается от Abstract Factory, тем, что Abstract Factory производит семейство объектов,
// эти объекты разные, обладают разными интерфейсами, но взаимодействуют между собой.
// В то время как Factory Method производит продукты придерживающиеся одного интерфейса
//  и эти продукты не связаны между собой, не вступают во взаимодействие.

package main

// each product should be usable
type Product interface {
	Use() string
}

// реализация интерфейса - конкретный продукт
type ProductA struct {
	action string
}

func (p ProductA) Use() string {
	return p.action
}

// реализация интерфейса - конкретный продукт
type ProductB struct {
	action string
}

func (p ProductB) Use() string {
	return p.action
}

// реализация интерфейса - конкретный продукт
type ProductC struct {
	action string
}

func (p ProductC) Use() string {
	return p.action
}

// интерфейс фабрики
type Factory interface {
	CreateProduct(Type string) Product
}

// конкретная фабрика
type FactoryA struct {
}

// конструктор фабрики
func NewFactory() *FactoryA {
	return &FactoryA{}
}

// фабричный метод
func (f *FactoryA) CreateProduct(Type string) Product {
	var product Product

	switch Type {
	case "A":
		product = ProductA{}
	case "B":
		product = ProductB{}
	case "C":
		product = ProductC{}
	}

	return product
}
