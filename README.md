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

#### Abstract Factory
An [Abstract Factory](./creational/factory/abstract/) is a factory of factories, it is, a factory that groups the different related/dependent factories together without specifying their concrete classes.

### Prototype
[Prototype](./creational/prototype/) specifies the kinds of objects to create using a prototypical instance, and creates new objects from the 'skeleton' of an existing object, thus boosting performance and keeping memory footprints to a minimum.


## Structural Design Patterns
Structural design patterns are design patterns that ease the design by identifying a simple way to realize relationships among entities.

### Composite
Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

### Adapter
Convert the interface of a class into another interface clients expect. An adapter lets classes work together that could not otherwise because of incompatible interfaces.

### Bridge
Decouple an abstraction from its implementation allowing the two to vary independently.

## Behavioral Design Patterns
Behavioral design patterns are design patterns that identify common communication patterns among objects and realize these patterns

## Concurrency Patterns
Concurrency patterns are those types of design patterns that deal with the multi-threaded programming paradigm. 