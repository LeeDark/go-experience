package main

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}
