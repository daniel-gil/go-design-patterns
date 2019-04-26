package travel_factory

// We have 2 entities
type Booking interface {
	GetBookingNumber() string
}
type Invoice interface {
	GetInvoiceNumber() string
}

type AbstractFactory interface {
	CreateBooking() Booking
	CreateInvoice() Invoice
}

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
