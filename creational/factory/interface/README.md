# Interface Factory

An `Interface Factory` returns interfaces instead of structs. This means that we can make structs private, while only exposing the interface outside our package.


```go
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
```