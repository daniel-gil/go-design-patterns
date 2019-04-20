# Simple Factory
The simplest factory is a function that takes in some arguments and returns an instance of a struct. 
It simulates the OOP constructor concept.

Suppose we have the following `Animal` struct, and that it has associated the `Eat` method:

```go
type Animal struct {
  Name string
  Kind string
}

func (a Animal) Eat() {
  fmt.Printf("The animal named %s is eating", a.Name)
}
```

Then we can create a factory for creating an instance of `Animal`:
```go
func NewAnimal(name, kind string) Animal {
  return Animal{
    Name: name,
    Kind: kind,
  }
}
```

The benefit of using factories instead of directly
```go
a := &Animal{}
```
is that the factory function signature ensures that the required attributes always will be supplied.
