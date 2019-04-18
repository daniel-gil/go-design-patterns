# Go Design Patterns
Go design patterns compilation grouped into the following categories: **Creational**, **Structural**, **Behavioral**, and **Concurrency**.


## Creational Design Patterns
Creational design patterns are design patterns that deal with object creation mechanisms, trying to create objects in a manner suitable to the situation. 

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


## Structural Design Patterns
Structural design patterns are design patterns that ease the design by identifying a simple way to realize relationships among entities.

## Behavioral Design Patterns
Behavioral design patterns are design patterns that identify common communication patterns among objects and realize these patterns

## Concurrency Patterns
Concurrency patterns are those types of design patterns that deal with the multi-threaded programming paradigm. 