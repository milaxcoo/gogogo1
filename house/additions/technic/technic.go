package technic

import (
	"fmt"
)

type Technic struct {
	Type string
	Color string
	Smart bool
}

func (t Technic) Print() {
	fmt.Print("Тип устройства: ", t.Type, "\nЦвет: ", t.Color, "\n")
	if t.Smart {
		fmt.Print("Умный дом: Да\n")
	} else {
		fmt.Print("Умный дом: Нет\n")
	}
}
