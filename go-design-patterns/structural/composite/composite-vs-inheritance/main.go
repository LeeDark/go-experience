package main

import "fmt"

type Parent struct {
	SomeField int
}

type Son1 struct {
	Parent
}

type Son2 struct {
	P Parent
}

func GetParentField(p Parent) int {
	fmt.Println(p.SomeField)
	return p.SomeField
}

func main() {
	son1 := Son1 {
		Parent{SomeField: 1},
	}

	fmt.Println("Use son1:", son1)

	// cannot use son (type Son) as type Parent in argument to GetParentField
	//fmt.Println("GetParentField:", GetParentField(son1))

	son2 := Son2 {
		P: Parent{SomeField: 2},
	}

	fmt.Println("Use son2:", son2)

	fmt.Println("GetParentField:", GetParentField(son2.P))
}