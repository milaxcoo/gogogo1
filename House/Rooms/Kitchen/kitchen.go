package kitchen

import (
	furniture "gogogo1/House/Furniture"
	technic "gogogo1/House/Technic"
	family "gogogo1/House/Family"
)


type  Kitchen struct{
	Table furniture.Furniture
	Chair furniture.Furniture
	Oven technic.Technic
	Family family.Family
}

//print func
func (kitchen Kitchen) String() string {
	return "Кухня: " + "\n" + kitchen.Table.String() + "\n" + kitchen.Chair.String() + "\n" + kitchen.Oven.String() + "\n" + kitchen.Family.String()
}


