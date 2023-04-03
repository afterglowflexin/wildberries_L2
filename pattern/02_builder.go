// определяет процесс поэтапного построения сложного продукта
package main

import "fmt"

type Builder interface {
	work1()
	work2()
	work3()
}

type Manager struct {
	builder Builder
}

func (m Manager) Construct() {
	m.builder.work1()
	m.builder.work2()
	m.builder.work3()
}

type ConcreteBuilder struct {
	product *Product
}

func (c ConcreteBuilder) work1() {
	c.product.content += "work1 done; "
}

func (c ConcreteBuilder) work2() {
	c.product.content += "work2 done; "
}

func (c ConcreteBuilder) work3() {
	c.product.content += "work3 done; "
}

type Product struct {
	content string
}

func (p Product) Show() {
	fmt.Println(p.content)
}

func main() {
	product := &Product{}
	Manager{ConcreteBuilder{product}}.Construct()
	product.Show()
}
