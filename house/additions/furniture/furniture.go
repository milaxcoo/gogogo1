package furniture

import (
	"fmt"
)

type Furniture struct {
	Type string
	Color string
	ComfortRating int
}

func (f Furniture) Print() {
	fmt.Print("Тип мебели: ", f.Type, "\nЦвет: ", f.Color, "\nРейтинг комфорта: ", f.ComfortRating, "\n")
}