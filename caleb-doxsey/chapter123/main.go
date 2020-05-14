package main

import "fmt"

// Scope
var globalVariable int = 100

var (
	globalVarA int = 1
	globalVarB float32 = 0.5
)

const globalConst int = 200

const (
	globalConstA int = 1
	globalConstB float32 = 0.5
)

// here can be a comment about function main
func main() {
	// Hello
	fmt.Println("Hello World")

	// Numbers: Integers
	var a int = 1
	var b int; b = 2
	var c int
	c = 3
	d := 4 	// var d int

	var e int16 = 5
	var f uint32 = 6
	var g byte = 7

	fmt.Println("Integers:", a, b, c, d, e, f, g)

	// Numbers: Floating Point Numbers
	var x float32 = 1.234
	var y float64 = 2.345678
	var z = 3.456789 	// var z float64

	fmt.Println("Floating:", x, y, z)

	var fa float32 = .1
	var fb float32 = .2
	fmt.Println(fa + fb)

	var fc float32 = .1
	var fd float32 = .2
	fmt.Println(fc + fd)

	// Strings
	var s string = "Hello World"
	fmt.Println(s, len(s))

	var t = "Another string"
	fmt.Println(t, len(t))

	fmt.Println("Hello World"[1]) 			// result will be [int]
	fmt.Println("Hello " + "World") 	// string concatenation

	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variants
	var sx string
	sx = "first "
	fmt.Println(sx)
	sx = sx + "second"
	fmt.Println(sx)

	var sx1 string = "hello"
	var sx2 string = "world"
	fmt.Println(sx1 == sx2)

	// About Naming
	n := "Max"
	fmt.Println("My dog's name is", n)
	name := "Max"
	fmt.Println("My dog's name is", name)
	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	// Constants
	const helloWorld string = "Hello World"
	fmt.Println(helloWorld)

	// Scope
	fmt.Println(globalVariable)
	fmt.Println(globalConst)

	fmt.Println(globalVarA)
	fmt.Println(globalVarB)

	fmt.Println(globalConstA)
	fmt.Println(globalConstB)
}