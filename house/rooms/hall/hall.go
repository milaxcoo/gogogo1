package hall

import (
	technic "gogogo1/house/technic"
	furniture "gogogo1/house/furniture"
	family "gogogo1/house/family"
)

type Hall struct {
	Tv technic.Technic
	Sofa furniture.Furniture
	Table furniture.Furniture
	Chair furniture.Furniture
	Phone technic.Technic
	Family family.Family
}

