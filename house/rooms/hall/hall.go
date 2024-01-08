package hall

import (
	"fmt"
	technic "gogogo1/house/additions/technic"
	furniture "gogogo1/house/additions/furniture"
)

type Hall struct {
	Tv technic.Technic
	Sofa furniture.Furniture
	Table furniture.Furniture
	Chair furniture.Furniture
	Phone technic.Technic
}

func (h Hall) Print() {
	fmt.Print("\nЗал:\n")
	h.Tv.Print()
	h.Sofa.Print()
	h.Table.Print()
	h.Chair.Print()
	h.Phone.Print()
}