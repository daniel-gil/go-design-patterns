# Abstract Factory
An `Abstract Factory` is a factory of factories, it is, a factory that groups the 
different related/dependent factories together without specifying their concrete classes.

Suppose we design a booking system for a travel agency where we have two generic 
entities: `Booking` and `Invoice`
```go
type Booking interface {
    GetBookingNumber() string
}
type Invoice interface {
    GetInvoiceNumber() string
}
```

and we create an abstract factory for creating those entities:

```go
type AbstractFactory interface {
    CreateBooking() Booking
    CreateInvoice() Invoice
}
```

Now we define two concrete product types, `hotel` and `flight`, where each one 
must implement the `AbstractFactory` using its own factory.

Here is the `HotelFactory`:
```go

const HotelPrefix = "HOT"

type HotelFactory struct{}

func (f HotelFactory) CreateBooking() Booking {
    return &HotelBooking{
        bookingNumber: fmt.Sprintf("%s-BKG-%s", HotelPrefix, time.Now().UnixNano()),
    }
}

func (f HotelFactory) CreateInvoice() Invoice {
    return &HotelInvoice{
        invoiceNumber: fmt.Sprintf("%s-INV-%s", HotelPrefix, time.Now().UnixNano()),
    }
}

type HotelBooking struct {
    bookingNumber string
}

func (b HotelBooking) GetBookingNumber() string {
    return b.bookingNumber
}

type HotelInvoice struct {
    invoiceNumber string
}

func (i HotelInvoice) GetInvoiceNumber() string {
    return i.invoiceNumber
}
```


and here is the `FlightFactory`:
```go
const FlightPrefix = "FLT"

type FlightFactory struct{}

func (f FlightFactory) CreateBooking() Booking {
    return &FlightBooking{
        bookingNumber: fmt.Sprintf("%s-BKG-%s", FlightPrefix, time.Now().UnixNano()),
    }
}

func (f FlightFactory) CreateInvoice() Invoice {
    return &FlightInvoice{
        invoiceNumber: fmt.Sprintf("%s-INV-%s", FlightPrefix, time.Now().UnixNano()),
    }
}

type FlightBooking struct {
    bookingNumber string
}

func (b FlightBooking) GetBookingNumber() string {
    return b.bookingNumber
}

type FlightInvoice struct {
    invoiceNumber string
}

func (i FlightInvoice) GetInvoiceNumber() string {
    return i.invoiceNumber
}
```

Finally we need to implement a function to get the proper factory depending 
on the type of product:
```go
type ProductType string

const (
    FLIGHT ProductType = "flight"
    HOTEL  ProductType = "hotel"
)

func GetFactory(productType ProductType) AbstractFactory {
    var factory AbstractFactory

    switch productType {
    case FLIGHT:
        factory = FlightFactory{}
    case HOTEL:
        factory = HotelFactory{}
    }

    return factory
}
```


## Usage
The client will use the abstract factory like this:

```go 
package main

import (
    "fmt"

    tf "github.com/daniel-gil/go-design-patterns/creational/factory/abstract/travel_factory"
)

func main() {
    hotelFactory := tf.GetFactory(tf.HOTEL)

    hotelBooking := hotelFactory.CreateBooking()
    fmt.Printf("Hotel Booking Number: %s\n", hotelBooking.GetBookingNumber())

    hotelInvoice := hotelFactory.CreateInvoice()
    fmt.Printf("Hotel Invoice Number: %s\n", hotelInvoice.GetInvoiceNumber())

    flightFactory := tf.GetFactory(tf.FLIGHT)

    flightBooking := flightFactory.CreateBooking()
    fmt.Printf("Flight Booking Number: %s\n", flightBooking.GetBookingNumber())

    flightInvoice := flightFactory.CreateInvoice()
    fmt.Printf("Flight Invoice Number: %s\n", flightInvoice.GetInvoiceNumber())
}
```


Output:
```bash
$ go run main.go 
Hotel Booking Number: HOT-BKG-1556278877443088000
Hotel Invoice Number: HOT-INV-1556278877443139000
Flight Booking Number: FLT-BKG-1556278877443144000
Flight Invoice Number: FLT-INV-1556278877443158000
```