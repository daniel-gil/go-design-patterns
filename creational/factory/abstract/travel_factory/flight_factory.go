package travel_factory

import (
	"fmt"
	"time"
)

const FlightPrefix = "FLT"

type FlightFactory struct{}

func (f FlightFactory) CreateBooking() Booking {
	return &FlightBooking{
		bookingNumber: fmt.Sprintf("%s-BKG-%v", FlightPrefix, time.Now().UnixNano()),
	}
}

func (f FlightFactory) CreateInvoice() Invoice {
	return &FlightInvoice{
		invoiceNumber: fmt.Sprintf("%s-INV-%v", FlightPrefix, time.Now().UnixNano()),
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
