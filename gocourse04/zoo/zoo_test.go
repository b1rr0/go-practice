package zoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZoo(t *testing.T) {
	z := NewZoo()
	area := *NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")

	z.Areas[AllAnimalSpeciesValues()[0]] = area

	area = *NewArea("Test area2", "Test area2")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector2")
	z.Areas[AllAnimalSpeciesValues()[1]] = area

	area = *NewArea("Test area3", "Test area3")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector3")
	z.Areas[AllAnimalSpeciesValues()[2]] = area

	lion := NewAnimal(AllAnimalSpeciesValues()[0], 1, "Lion")
	z.AddAnimal(lion)
	leopard := NewAnimal(AllAnimalSpeciesValues()[1], 2, "Leopard")
	z.AddAnimal(leopard)
	elephant := NewAnimal(AllAnimalSpeciesValues()[2], 3, "Elephant")
	z.AddAnimal(elephant)

	assert.Equal(t, lion.Name, z.GetAnimalByName(lion.Name).Name)
	assert.Equal(t, leopard.Name, z.GetAnimalByName(leopard.Name).Name)
	assert.Equal(t, elephant.Name, z.GetAnimalByName(elephant.Name).Name)

	assert.True(t, z.MoveAnimal(lion.Name, SectorTypeAviary))
}

func TestZooAddAnimalToNonExistingArea(t *testing.T) {
	z := NewZoo()
	area := *NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")
	z.Areas[AllAnimalSpeciesValues()[0]] = area

	_, err := z.AddAnimal(NewAnimal(AllAnimalSpeciesValues()[0], 1, "Lion"))
	assert.NoError(t, err)

	_, err = z.AddAnimal(NewAnimal(AllAnimalSpeciesValues()[1], 1, "Leopard"))
	assert.Error(t, err)
}

func TestZooGetAnimalByNameNonExisting(t *testing.T) {
	z := NewZoo()
	area := *NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")
	z.Areas[AllAnimalSpeciesValues()[0]] = area
	assert.Nil(t, z.GetAnimalByName("Lion"))
}

func TestZooMoveAnimalNonExisting(t *testing.T) {
	z := NewZoo()

	assert.False(t, z.MoveAnimal("Lion", SectorTypeAviary))
}

func TestZooMoveAnimalSector(t *testing.T) {
	z := NewZoo()
	area := *NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")
	z.Areas[AllAnimalSpeciesValues()[0]] = area

	z.AddAnimal(NewAnimal(AllAnimalSpeciesValues()[0], 1, "Lion"))
	assert.True(t, z.MoveAnimal("Lion", SectorTypeCage))
}

func BenchmarkGetAnimalByName(b *testing.B) {
	z := NewZoo()
	area := *NewArea("Test area", "Test area")
	area.Sectors[DefaultSectorType] = *NewSector("Test sector")
	z.Areas[AllAnimalSpeciesValues()[0]] = area

	z.AddAnimal(NewAnimal(AllAnimalSpeciesValues()[0], 1, "Lion"))

	b.ResetTimer() // Скидаємо таймер перед початком тестування

	for i := 0; i < b.N; i++ {
		_ = z.GetAnimalByName("Lion")
	}
}
