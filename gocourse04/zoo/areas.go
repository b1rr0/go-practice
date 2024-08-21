package zoo

import "fmt"

type Areas map[AnimalSpecies]Area

type Area struct {
	Name    string
	Type    string
	Sectors map[SectorType]Sector
}

func NewArea(name string, areaType string) *Area {
	return &Area{
		Name:    name,
		Type:    areaType,
		Sectors: make(map[SectorType]Sector),
	}
}

func (a Area) Show() {
	fmt.Print("Area+", a.Name, "  \n", a.Type, "\n")
	for _, sector := range a.Sectors {
		sector.Show()
	}
	fmt.Print("==================\n")
}

func (a Area) AddAnimalToDefaultSector(animal *Animal) {
	a.Sectors[DefaultSectorType] = a.Sectors[DefaultSectorType].AddAnimal(animal)
}

func (a Area) AddAnimalToSector(animal *Animal, sectorType SectorType) {
	a.Sectors[sectorType] = a.Sectors[sectorType].AddAnimal(animal)
}
