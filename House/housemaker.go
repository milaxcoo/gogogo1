package house

import (
	"fmt"
	family "gogogo1/House/Family"
	furniture "gogogo1/House/Furniture"
	bedroom "gogogo1/House/Rooms/Bedroom"
	hall "gogogo1/House/Rooms/Hall"
	kitchen "gogogo1/House/Rooms/Kitchen"
	technic "gogogo1/House/Technic"
)

type House struct{
	Hall hall.Hall
	Kitchen kitchen.Kitchen
	Bedroom bedroom.Bedroom
}


var table = furniture.Furniture{
	Name: 		"Стол IKEA",
	Color: 		"Белый",
	Price: 		1000,
}
var chair = furniture.Furniture{
	Name : 		"Стул IKEA",
	Color: 		"Белый",
	Price: 		500,
}
var bed = furniture.Furniture{
	Name : 		"Кровать Saatva",
	Color: 		"Коричневый",
	Price: 		10000,
}
var wardrobe = furniture.Furniture{
	Name : 		"Шкаф IKEA",
	Color: 		"Белый",
	Price: 		5000,
}
var tv = technic.Technic{
	Name: 		"Телевизор Samsung",
	Color: 		"Черный",
	Price: 		10000,
}
var sofa = furniture.Furniture{
	Name : 		"Диван IKEA",
	Color: 		"Белый",
	Price: 		10000,
}
var phone = technic.Technic{
	Name: 		"Телефон Samsung",
	Color: 		"Черный",
	Price: 		10000,
}
var radio = technic.Technic{
	Name: 		"Радио Sony",
	Color: 		"Черный",
	Price: 		10000,
}
var oven = technic.Technic{
	Name: 		"Духовка Samsung",
	Color: 		"Черный",
	Price: 		10000,
}
var dad = family.Family{
	Sex: 		"Мужчина",
	Name: 		"Василий",
	Age: 		45,
	Guest: 		false,
}
var mom = family.Family{
	Sex: 		"Женщина",
	Name: 		"Анна",
	Age: 		40,
	Guest: 		false,
}
var son = family.Family{
	Sex: 		"Мужчина",
	Name: 		"Алексей",
	Age: 		20,
	Guest: 		false,
}
var guest = family.Family{
	Sex: 		"Девушка",
	Name: 		"Алиса",
	Age: 		19,
	Guest: 		true,
}

func housemaker() House{
	var house House
	house.Hall = hall.Hall{
		Tv: tv,
		Sofa: sofa,
		Table: table,
		Chair: chair,
		Phone: phone,
		Family: dad,
	}
	house.Kitchen = kitchen.Kitchen{
		Table: table,
		Chair: chair,
		Oven: oven,
		Family: mom,
	}
	house.Bedroom = bedroom.Bedroom{
		Table: table,
		Chair: chair,
		Bed: bed,
		Wardrobe: wardrobe,
		Family: son,
		Radio: radio,
	}
	fmt.Println("Дом собран")
	return house
}

