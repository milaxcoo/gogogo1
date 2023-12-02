package family


type Family struct {
	Sex string
	Age int
	Name string
	Guest bool
}

//print func
func (family Family) String() string {
	return "Живут: " + "\n" + family.Name + "\n"
}