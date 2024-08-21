package zoo

import "fmt"

type Animal struct {
	Species AnimalSpecies
	ID      int
	Name    string
}

func NewAnimal(species AnimalSpecies, id int, name string) *Animal {
	return &Animal{
		Species: species,
		ID:      id,
		Name:    name,
	}
}

func (a Animal) Show() {
	fmt.Println("Animal name:->" + a.Name + " <->" + a.Species.String())
}
