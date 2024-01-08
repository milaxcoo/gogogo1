package familyliving

import (
	"fmt"
)

type Family struct {
	Sex bool
	Age int
	Name string
	Status bool
	Children int
}

func (f Family) Print() {
	fmt.Print("Имя: ", f.Name, "\nВозраст: ", f.Age, "\n")
	if f.Sex {
		fmt.Print("Пол: Мужской\n")
		if f.Status {
			fmt.Print("Семейное положение: Женат\n")
		} else {
			fmt.Print("Семейное положение: Не женат\n")
		}
	} else {
		fmt.Print("Пол: Женский\n")
		if f.Status {
			fmt.Print("Семейное положение: Замужем\n")
		} else {
			fmt.Print("Семейное положение: Не замужем\n")
		}
	}
	if f.Children == 0 {
		fmt.Print("Дети: Нет\n")
	} else {
		fmt.Print("Дети: ", f.Children, "\n")
	}

}