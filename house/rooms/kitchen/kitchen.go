package kitchen

import (
	furniture "gogogo1/house/furniture"
	technic "gogogo1/house/technic"
	family "gogogo1/house/family"
)


type  Kitchen struct{
	Table furniture.Furniture
	Chair furniture.Furniture
	Oven technic.Technic
	Family family.Family
}
