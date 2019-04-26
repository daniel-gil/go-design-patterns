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
