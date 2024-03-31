package main

// Create the struc or type or the blueprint
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make New bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
