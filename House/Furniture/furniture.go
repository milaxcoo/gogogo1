package furniture

type Furniture struct {
	Name string
	Color string
	Price int
}

//print func
func (furniture Furniture) String() string {
	return "Мебель: " + "\n" + "Название: " + furniture.Name + "\n" + "Цвет: " + furniture.Color + "\n" + "Цена: " + string(furniture.Price) + "\n"
}