package familyliving

import (
	"fmt"
)

type FamilyLiving struct {
	Sex bool
	Name string
	Age int
	Guest bool
	Status bool
}

func (f FamilyLiving) Print() {
	fmt.Print("\nИмя: ", f.Name, "\nВозраст: ", f.Age, "\n")
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
	
}