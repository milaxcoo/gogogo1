package bedroom

import (
	furniture "gogogo1/house/additions/furniture"
	technic "gogogo1/house/additions/technic"
	family "gogogo1/house/additions/family"
	family_living "gogogo1/house/additions/family/family_living"

)


type Bedroom struct {
	Table furniture.Furniture
	Chair furniture.Furniture
	Bed furniture.Furniture
	Wardrobe furniture.Furniture
	Family family.Family
	Radio technic.Technic
}

func (b Bedroom) Print() {
	b.Family.Print()
	b.Table.Print()
	b.Chair.Print()
	b.Bed.Print()
	b.Wardrobe.Print()
	b.Radio.Print()
}

