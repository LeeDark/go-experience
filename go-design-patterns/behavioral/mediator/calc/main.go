package main

import "fmt"

type One struct{}
type Two struct{}
type Three struct{}
type Four struct{}

func Sum(a, b interface{}) interface{} {
	switch a := a.(type) {
	case One:
		switch b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		default:
			return fmt.Errorf("Number not found")
		}
	case Two:
		switch b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		default:
			return fmt.Errorf("Number not found")
		}
	case int:
		switch b := b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		case int:
			return a + b
		default:
			return fmt.Errorf("Number not found")
		}
	default:
		return fmt.Errorf("Number not found")
	}
}

func main() {
	fmt.Printf("%#v\n", Sum(One{}, Two{}))
	fmt.Printf("%d\n", Sum(1, 2))
}
