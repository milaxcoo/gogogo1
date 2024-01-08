package family

import (
	"fmt"
	family_living "gogogo1/house/additions/family/family_living"
)

type Family struct {
	Family []family_living.Family
	Surname string
}

func (f Family) Print() {
	fmt.Print("Семья: ", f.Surname, "\n")
	for _, family := range f.Family {
		family.Print()
	}
}