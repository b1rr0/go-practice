package zoo

import "fmt"

type Zoo struct {
	Areas      Areas
	SetOfNames map[string]struct{} // smthiong like bloom filter
}

func NewZoo() *Zoo {
	return &Zoo{
		Areas:      make(map[AnimalSpecies]Area),
		SetOfNames: make(map[string]struct{}),
	}
}

func (z *Zoo) AddAnimal(animal *Animal) (bool, error) {
	if _, exists := z.SetOfNames[animal.Name]; exists {
		return false, fmt.Errorf("Animal %s already exists", animal.Name)
	}
	z.SetOfNames[animal.Name] = struct{}{}

	area, exists := z.Areas[animal.Species]
	if !exists {
		return false, fmt.Errorf("Area for %s does not exist", animal.Species)
	}

	area.AddAnimalToDefaultSector(animal)

	return true, nil
}

func (z *Zoo) GetAnimalByName(name string) *Animal {
	if _, exists := z.SetOfNames[name]; !exists {
		return nil
	}

	for _, area := range z.Areas {
		for _, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					return &animal
				}
			}
		}
	}
	return nil
}

func (z *Zoo) Show() {
	for name, area := range z.Areas {
		fmt.Print(name.String() + "  contains:\n")
		area.Show()
	}
}

func (z *Zoo) MoveAnimal(name string, sectorType SectorType) bool {
	sID, s := z.lookup(name)
	if s == nil {
		return false
	}

	animal := z.GetAnimalByName(name)
	fmt.Print(animal)

	// transaction
	z.Areas[animal.Species].Sectors[sID] = s.RemoveAnimal(name)
	z.Areas[animal.Species].AddAnimalToSector(animal, sectorType)

	return true
}

func (z *Zoo) lookup(name string) (SectorType, *Sector) {
	if _, exists := z.SetOfNames[name]; !exists {
		return SectorTypeAviary, nil
	}

	for _, area := range z.Areas {
		for id, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					return id, &sector
				}
			}
		}
	}
	return SectorTypeAviary, nil
}

func (z *Zoo) Lookup(name string) *Sector {
	_, sector := z.lookup(name)
	return sector
}
