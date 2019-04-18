package factory

import "fmt"

type Animal interface {
	Eat()
}

type animal struct {
	name string
	kind string
}

func (a animal) Eat() {
	fmt.Printf("The animal named %s is eating", a.name)
}

func NewAnimal(name, kind string) Animal {
	return &animal{
		name: name,
		kind: kind,
	}
}
