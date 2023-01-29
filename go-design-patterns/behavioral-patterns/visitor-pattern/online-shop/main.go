package main

import (
	"fmt"

	"github.com/LeeDark/go-experience/go-design-patterns/behavioral-patterns/visitor-pattern/online-shop/visitor"
)

func main() {
	products := make([]visitor.Visitable, 2)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}

	// Print the sum of prices
	priceVisitor := &PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	// Print the product list
	nameVisitor := &NamePrinter{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n---------------\n%s", nameVisitor.ProductList)
}