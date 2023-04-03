// visitor может обойти все разнородные elements в какой структуре(ObjectStructure)
package main

// visitor interface. visitor делает работу с разными элементами
type Visitor interface {
	VisitElementA(a *ElementA) string
	VisitElementB(b *ElementB) string
}

// конкретный visitor
type VisitorA struct {
}

// реализация интерфейса
func (v *VisitorA) VisitElementA(a *ElementA) string {
	return a.BuyA()
}

// реализация интерфейса
func (v *VisitorA) VisitElementB(b *ElementB) string {
	return b.BuyB()
}

// структура элементов
type ObjectStructure struct {
	elements []Element
}

// функция для обхода всех элементов посетителем
func (o *ObjectStructure) Accept(v Visitor) string {
	var result string
	for _, p := range o.elements {
		result += p.AcceptVisit(v)
	}
	return result
}

type Element interface {
	AcceptVisit(v Visitor) string
}

type ElementA struct {
}

func (a *ElementA) AcceptVisit(v Visitor) string {
	return v.VisitElementA(a)
}

func (a *ElementA) BuyA() string {
	return "Buy A"
}

type ElementB struct {
}

func (b *ElementB) AcceptVisit(v Visitor) string {
	return v.VisitElementB(b)
}

func (b *ElementB) BuyB() string {
	return "Buy B"
}
