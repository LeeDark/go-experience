package main

import (
	"fmt"

	"github.com/LeeDark/go-experience/go-design-patterns/behavioral-patterns/visitor-pattern/online-shop/visitor"
)

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p visitor.ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p visitor.ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
