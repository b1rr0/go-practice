package zoo

import "fmt"

type Sector struct {
	Subtype string
	Animals []Animal
}

func NewSector(subtype string) *Sector {
	return &Sector{
		Subtype: subtype,
		Animals: make([]Animal, 0),
	}
}

func (s Sector) AddAnimal(animal *Animal) Sector {
	s.Animals = append(s.Animals, *animal)
	return s
}

func (s Sector) Show() {
	fmt.Print("Sector+", s.Subtype, "  \n")
	for _, animal := range s.Animals {
		animal.Show()
	}
	fmt.Print("==================\n")
}

func (s Sector) RemoveAnimal(name string) Sector {
	for i, animal := range s.Animals {
		if animal.Name == name {
			s.Animals = append(s.Animals[:i], s.Animals[i+1:]...)
			break
		}
	}
	return s
}
