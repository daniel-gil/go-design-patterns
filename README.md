# Go Design Patterns
Go design patterns compilation based on the Udemy course [Go: Concurrency & Design Patterns for Gophers](https://www.udemy.com/learning-pathgo-concurrency-and-design-patterns-for-gophers/).

## Creational Patterns
### Singleton
[Singleton](./creational/singleton/) provides a single instance of an object and guarantees that there are no duplicates.

### Builder
[Builder](./creational/builder/) is a separate component for building a complex object.


### Factory
A `Factory` is a component responsible solely for the wholesale (not piecewise) creation of objects. 

#### Simple Factory
The [Simple Factory](./creational/factory/simple/) is a function that takes in some arguments and returns an instance of a struct.

#### Interface Factory
An [Interface Factory](./creational/factory/interface/)  returns interfaces instead of structs. This means we can make structs private, while only exposing the interface outside our package.

#### Factory Method

The factory method pattern defines an interface for creating a single object, but let subclasses decide which class to instantiate. [Factory Method](./creational/factory/method/) lets a class defer instantiation to subclasses.


### Prototype
[Prototype](./creational/prototype/) specifies the kinds of objects to create using a prototypical instance, and creates new objects from the 'skeleton' of an existing object, thus boosting performance and keeping memory footprints to a minimum.