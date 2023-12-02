package bedroom

import (
	furniture "gogogo1/House/Furniture"
	technic "gogogo1/House/Technic"
	family "gogogo1/House/Family"
)


type Bedroom struct {
	Table furniture.Furniture
	Chair furniture.Furniture
	Bed furniture.Furniture
	Wardrobe furniture.Furniture
	Family family.Family
	Radio technic.Technic
}

//print func
func (bedroom Bedroom) String() string {
	return "Спальня: " + "\n" + bedroom.Table.String() + "\n" + bedroom.Chair.String() + "\n" + bedroom.Bed.String() + "\n" + bedroom.Wardrobe.String() + "\n" + bedroom.Family.String() + "\n" + bedroom.Radio.String()
}