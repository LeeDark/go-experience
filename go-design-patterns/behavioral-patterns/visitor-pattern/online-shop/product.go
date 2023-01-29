package main

import "github.com/LeeDark/go-experience/go-design-patterns/behavioral-patterns/visitor-pattern/online-shop/visitor"

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v visitor.Visitor) {
	v.Visit(p)
}
