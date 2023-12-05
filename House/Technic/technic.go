package technic

import "fmt"

type Technic struct {
	Name string
	Color string
	Price int
}


func (technic Technic) String() string {
	return "Техника: " + "\n" + "Название: " + technic.Name + "\n" + "Цвет: " + technic.Color + "\n" + "Цена: " + fmt.Sprint(technic.Price) + "\n"
}
