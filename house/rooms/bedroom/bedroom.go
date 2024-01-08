package bedroom

import (
	"fmt"
	furniture "gogogo1/house/additions/furniture"
	technic "gogogo1/house/additions/technic"
)

type Bedroom struct {
	Table furniture.Furniture
	Chair furniture.Furniture
	Bed furniture.Furniture
	Wardrobe furniture.Furniture
	Radio technic.Technic
}

func (b Bedroom) Print() {
	fmt.Print("\nСпальня:\n")
	b.Table.Print()
	b.Chair.Print()
	b.Bed.Print()
	b.Wardrobe.Print()
	b.Radio.Print()
}