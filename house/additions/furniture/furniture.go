package furniture

import (
	"fmt"
)

type Furniture struct {
	Name string
	Color string
	Price int
}

func (f Furniture) Print() {
	fmt.Print("\nМебель: ", f.Name, "\nЦвет: ", f.Color, "\nЦена: ", f.Price, "\n")
}