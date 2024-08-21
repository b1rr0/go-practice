package main

import (
	"gocourse04/zoo"
	"strconv"
)

func main() {
	z := BuildZoo()
	FeelWithRandomAnimals(z)
	z.Show()

	z.MoveAnimal("Monkey10", zoo.SectorTypeAviary)
	z.MoveAnimal("Monkey9", zoo.SectorTypeAviary)
	z.MoveAnimal("Monkey8", zoo.SectorTypeAviary)
	z.MoveAnimal("Monkey7", zoo.SectorTypeAviary)
	//z.MoveAnimal("Monkey6", zoo.SectorTypeAviary)

	z.Show()
}

func FeelWithRandomAnimals(z *zoo.Zoo) {
	for _, animalSpecies := range zoo.AllAnimalSpeciesValues() {
		for i := 1; i <= 10; i++ {
			z.AddAnimal(zoo.NewAnimal(animalSpecies, 1, animalSpecies.String()+strconv.Itoa(i)))
		}
	}
}
func BuildZoo() *zoo.Zoo {
	z := zoo.NewZoo()

	for _, animalSpecies := range zoo.AllAnimalSpeciesValues() {
		area := *zoo.NewArea("Name + "+animalSpecies.String(), "Area: "+animalSpecies.String())
		for _, sectorType := range zoo.AllSectorTypeValues() {
			area.Sectors[sectorType] = *zoo.NewSector(sectorType.String() + " sector ")
		}

		z.Areas[animalSpecies] = area
	}

	return z
}
