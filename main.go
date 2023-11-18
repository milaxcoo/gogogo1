package main

import (
	"fmt"
	furniture "gogogo1/House/Furniture"
	
)



func main() {
	// Create a new sofa with the name "Sofa" and the color "Red".
	sofa := furniture.Sofa{
		Name:  "Sofa",
		Color: "Red",
		Price: 100,
	}

	// Print the name of the sofa.
	fmt.Println(sofa.Name)

	// Print the color of the sofa.
	fmt.Println(sofa.Color)

	// Print the price of the sofa.
	fmt.Println(sofa.Price)
}