package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

// 1: composite without embed
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

// 2: composite with embed
type CompositeShark struct {
	Animal
	Swim func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (si *SwimmerImpl) Swim(){
	fmt.Println("Swimming...")
}

// 3: composite with interfaces
type CompositeSwimmerB struct{
	Trainer
	Swimmer
}

func main() {
	localSwim := Swim

	swimmerA := CompositeSwimmerA{
		MySwim: localSwim,
	}

	swimmerA.MyAthlete.Train()
	swimmerA.MySwim()

	fish := CompositeShark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()
}
