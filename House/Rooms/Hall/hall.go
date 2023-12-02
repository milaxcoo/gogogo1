package hall

import (

	technic "gogogo1/House/Technic"
	furniture "gogogo1/House/Furniture"
	family "gogogo1/House/Family"
)

type Hall struct {
	Tv technic.Technic
	Sofa furniture.Furniture
	Table furniture.Furniture
	Chair furniture.Furniture
	Phone technic.Technic
	Family family.Family
}

//print func
func (hall Hall) String() string {
	return "Холл: " + "\n" + hall.Tv.String() + "\n" + hall.Sofa.String() + "\n" + hall.Table.String() + "\n" + hall.Chair.String() + "\n" + hall.Phone.String() + "\n" + hall.Family.String()
}