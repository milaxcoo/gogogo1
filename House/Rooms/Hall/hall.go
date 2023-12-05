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

