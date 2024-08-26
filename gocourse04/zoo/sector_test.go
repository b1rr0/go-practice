package zoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSector(t *testing.T) {
	sector := NewSector("Test sector")

	assert.Equal(t, "Test sector", sector.Subtype)
	assert.Equal(t, 0, len(sector.Animals))
}

func TestSectorAddAnimal(t *testing.T) {
	sector := *NewSector("Test sector")
	animal := NewAnimal(AllAnimalSpeciesValues()[0], 5, "qwe")

	sector = sector.AddAnimal(animal)

	assert.Equal(t, 1, len(sector.Animals))
	assert.Equal(t, 5, sector.Animals[0].ID)
}

func TestSectorRemoveAnimal(t *testing.T) {
	sector := *NewSector("Test sector")
	animal1 := NewAnimal(AllAnimalSpeciesValues()[0], 5, "qwe")
	animal2 := NewAnimal(AllAnimalSpeciesValues()[1], 6, "asd")
	sector = sector.AddAnimal(animal1)
	sector = sector.AddAnimal(animal2)

	sector = sector.RemoveAnimal("qwe")

	assert.Equal(t, 1, len(sector.Animals))
	assert.Equal(t, 6, sector.Animals[0].ID)
}
