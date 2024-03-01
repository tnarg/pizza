package parlor

import (
	"github.com/alecthomas/repr"
)

type Crust int

const (
	Wheat Crust = iota
	Cornmeal
)

type Topping int

const (
	Pepperoni Topping = iota
	Sausage
	Mushrooms
)

type Pizza struct {
	crust    Crust
	toppings []Topping
}

func (p *Pizza) String() string {
	return repr.String(p)
}

func NewPizza(crust Crust, toppings []Topping) *Pizza {
	return &Pizza{crust, toppings}
}
