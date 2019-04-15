# Builder

## Definition
The builder separates the construction of a complex object from its representation, allowing the same construction process to create various representations.

It helps to construct complex objects without directly instantiating their struct. Creates an object step by step by filling its fields and creating embedded objects. Avoid writing the logic to create all the objects in the package.

## Components
The builder pattern has the following components:

- `Director`: Director performs all the required steps given inside builder to create a product. It will invoke all the abstract builder steps to create our vehicle.

- `Abstract Builder`: Abstract builder provides common interface for concrete builder. All the method inside our abstract builder will be implemented by our concrete vehicle builders.

- `Concrete Builder`: Implements the builder interface and provides an interface for getting the product. There will be multiple concrete builders, one for each kind of product. In our case `car`, `bike` and `motorbike` are the concrete builders.

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

// we could override the default values to fit some specific needs
car.Color = RED
```

## Example
The file `builder.go` includes an implementation of the builder pattern.

## Test

## References
- [Builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) definition from Wikipedia