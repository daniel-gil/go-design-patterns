package factory

import "fmt"

type Animal struct {
	Name string
	Kind string
}

func (a Animal) Eat() {
	fmt.Printf("The animal named %s is eating", a.Name)
}

func NewAnimal(name, kind string) Animal {
	return Animal{
		Name: name,
		Kind: kind,
	}
}
