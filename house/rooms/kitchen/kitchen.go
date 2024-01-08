package kitchen

import (
	"fmt"
	furniture "gogogo1/house/additions/furniture"
	technic "gogogo1/house/additions/technic"
)


type  Kitchen struct{
	Table furniture.Furniture
	Chair furniture.Furniture
	Oven technic.Technic
}

func (k Kitchen) Print() {
	fmt.Print("\nКухня:\n")
	k.Table.Print()
	k.Chair.Print()
	k.Oven.Print()
}
