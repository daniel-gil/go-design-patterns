# Builder

## Definition
The builder separates the construction of a complex object from its representation, allowing the same construction process to create various representations.

It helps to construct complex objects without directly instantiating their struct. Creates an object step by step by filling its fields and creating embedded objects. Avoid writing the logic to create all the objects in the package.

## Components
The builder pattern has the following components:

- `Director`: Director performs all the required steps given inside builder to create a product. It will invoke all the abstract builder steps to create our vehicle.

- `Abstract Builder`: Abstract builder provides common interface for concrete builder. All the method inside our abstract builder will be implemented by our concrete vehicle builders.

- `Concrete Builder`: Implements the builder interface and provides an interface for getting the product. There will be multiple concrete builders, one for each kind of product. In our case `car` and `motorbike` are the concrete builders.

- `Product`: is the main object that’s constructed. Represents the complex object under construction. In our case, vehicle is our product.

## Usage

```go
// create a director
director := ManufacturingDirector{}

// create an concrete builder
carBuilder := &CarBuilder{}

// pass the concrete builder to the director
director.SetBuilder(carBuilder)

// request the director to construct a concrete vehicle object
director.Construct()

// retrieve the car vehicle object
car := carBuilder.GetVehicle()
```

### Chaining
Once we have the `car` vehicle object, we could override the default values to fit some specific needs using the chaining method:

```go 
car.
    SetColor(RED).
    SetTopSpeed(150).
    SetDefaultSeats(2)
```

## Example
The file `director.go` contains the implementation of the director component.

The file `builder.go` contains the `BuildProcess` interface which corresponds to the abstract builder component.

The files `car.go` and `motorbike.go` contain the implementation of the concrete builder component for each vehicle type.

The file `product.go` contains the definition of the product component, it is the `VehicleProduct` structure.

## Test

```bash
$ cd creational/builder/
$ go test -v
=== RUN   TestCar
--- PASS: TestCar (0.00s)
=== RUN   TestMotorbike
--- PASS: TestMotorbike (0.00s)
PASS
ok      github.com/daniel-gil/go-design-patterns/creational/builder     0.009s
```

## References
- [Builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) definition from Wikipedia
- [Desing Patterns in Golang: Builder](http://blog.ralch.com/tutorial/design-patterns/golang-builder/) by Svetlin Ralchev.