package house

import (
	"fmt"
	"gogogo1/house/additions/family"
	"gogogo1/house/additions/furniture"
	"gogogo1/house/additions/technic"
	family_living "gogogo1/house/additions/family/family_living"
	"gogogo1/house/rooms/bedroom"
	"gogogo1/house/rooms/hall"
	"gogogo1/house/rooms/kitchen"


)

type House struct {
	Family        []family.Family
	Technic       []technic.Technic
	Furniture     []furniture.Furniture
	Kitchen 	 	kitchen.Kitchen
	Hall 		 	hall.Hall
	Bedroom 	 	bedroom.Bedroom
}

func (h House) Print() {
	fmt.Print("Дом: 2-х комнатная квартира\n")
	for _, families := range h.Family {
		families.Print()
	}
	h.Kitchen.Print()
	h.Hall.Print()
	h.Bedroom.Print()

}

func CreateHouse() House {
	var table = furniture.Furniture{
		Name: 		"Стол IKEA",
		Color: 		"Белый",
		Price: 		1000,
	}
	var chair = furniture.Furniture{
		Name : 		"Стулья Hoff",
		Color: 		"Серый",
		Price: 		500,
	}
	var bed = furniture.Furniture{
		Name : 		"Кровать Saatva",
		Color: 		"Коричневый",
		Price: 		10000,
	}
	var wardrobe = furniture.Furniture{
		Name : 		"Шкаф ПАКС",
		Color: 		"Бежевый",
		Price: 		5000,
	}
	var tv = technic.Technic{
		Name: 		"Телевизор Samsung",
		Color: 		"Желтый",
		Smart: 		true,
	}
	var sofa = furniture.Furniture{
		Name : 		"Диван Maiden",
		Color: 		"Пурпурный",
		Price: 		10000,
	}
	var phone = technic.Technic{
		Name: 		"Смартфон iPhone",
		Color: 		"Титановый",
		Smart: 		true,
	}
	var radio = technic.Technic{
		Name: 		"Радио DEXP",
		Color: 		"Синий",
		Smart: 		false,
	}
	var oven = technic.Technic{
		Name: 		"Духовка Xiaomi",
		Color: 		"Черный",
		Smart: 		true,
	}
	var dad = family_living.FamilyLiving{
		Sex: 		true,
		Name: 		"Дмитрий",
		Age: 		58,
		Guest: 		false,
		Status: 	true,
	}
	var mom = family_living.FamilyLiving{
		Sex: 		false,
		Name: 		"Светлана",
		Age: 		58,
		Guest: 		false,
		Status: 	true,
	}
	var son = family_living.FamilyLiving{
		Sex: 		true,
		Name: 		"Илья",
		Age: 		20,
		Guest: 		false,
		Status: 	false,
	}
	var guest = family_living.FamilyLiving{
		Sex: 		false,
		Name: 		"Алиса",
		Age: 		19,
		Guest: 		true,
		Status: 	false,
	}

	house := House{
		Family: []family.Family{family.Family{
			Family: []family_living.FamilyLiving{dad, mom, son, guest},
			Surname:       "Медведевы",
			},
		},
		Furniture:     []furniture.Furniture{table, chair, wardrobe, bed, sofa},
		Technic:       []technic.Technic{tv, radio, oven, phone},
		Kitchen: 	   kitchen.Kitchen{Table: table, Chair: chair, Oven: oven},
		Hall: 		   hall.Hall{Tv: tv, Sofa: sofa, Table: table, Chair: chair, Phone: phone},
		Bedroom: 	   bedroom.Bedroom{Table: table, Chair: chair, Bed: bed, Wardrobe: wardrobe, Radio: radio},
	}
	
	return house
}

