package travel_factory

import (
	"fmt"
	"time"
)

const HotelPrefix = "HOT"

type HotelFactory struct{}

func (f HotelFactory) CreateBooking() Booking {
	return &HotelBooking{
		bookingNumber: fmt.Sprintf("%s-BKG-%v", HotelPrefix, time.Now().UnixNano()),
	}
}

func (f HotelFactory) CreateInvoice() Invoice {
	return &HotelInvoice{
		invoiceNumber: fmt.Sprintf("%s-INV-%v", HotelPrefix, time.Now().UnixNano()),
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
