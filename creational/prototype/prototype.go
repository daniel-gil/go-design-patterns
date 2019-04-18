package prototype

// Product interface
type Product interface {
	SKU() string
	Name() string
	Clone() Product
}

type concreteProduct struct {
	sku  string
	name string
}

// SKU returns the SKU of the concreteProduct.
func (p *concreteProduct) SKU() string {
	return p.sku
}

// Name returns the name of the concreteProduct.
func (p *concreteProduct) Name() string {
	return p.name
}

// Clone creates a cloned new instance of a concreteProduct.
func (p *concreteProduct) Clone() Product {
	return &concreteProduct{
		sku:  p.sku,
		name: p.name,
	}
}
