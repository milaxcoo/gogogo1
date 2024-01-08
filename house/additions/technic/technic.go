package technic

import (
	"fmt"
)

type Technic struct {
	Name string
	Color string
	Smart bool
}

func (t Technic) Print() {
	fmt.Print("\nТехника: ", t.Name, "\nЦвет: ", t.Color, "\n")
	if t.Smart {
		fmt.Print("Умный дом: Да\n")
	} else {
		fmt.Print("Умный дом: Нет\n")
	}
}
