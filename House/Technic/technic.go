package technic

type Technic struct {
	Name string
	Color string
	Price int
}

//print func
func (technic Technic) String() string {
	return "Техника: " + "\n" + "Название: " + technic.Name + "\n" + "Цвет: " + technic.Color + "\n" + "Цена: " + string(technic.Price) + "\n"
}
