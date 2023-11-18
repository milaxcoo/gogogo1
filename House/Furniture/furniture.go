package furniture

import "fmt"

func init() {
	fmt.Println("Furniture package initialized")
}

// Furniture is a struct that represents a piece of furniture.
type Sofa struct {
	// Name is the name of the piece of furniture.
	Name string
	// Color is the color of the piece of furniture.
	Color string
	// Price is the price of the piece of furniture.
	Price int
}