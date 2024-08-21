package zoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaAddAnimalToDefaultSector(t *testing.T) {
	area := NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")
	animal := NewAnimal(AllAnimalSpeciesValues()[0], 5, "qwe")
	area.AddAnimalToDefaultSector(animal)

	assert.Equal(t, 5, area.Sectors[DefaultSectorType].Animals[0].ID)
}

func TestAreaAddAnimalToSector(t *testing.T) {
	area := NewArea("Test area", "Test area")
	area.Sectors[SectorTypeCage] = *NewSector("Test sector")
	animal := NewAnimal(AllAnimalSpeciesValues()[0], 5, "qwe")
	area.AddAnimalToSector(animal, SectorTypeCage)

	assert.Equal(t, 5, area.Sectors[SectorTypeCage].Animals[0].ID)
}
