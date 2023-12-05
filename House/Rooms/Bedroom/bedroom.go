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

