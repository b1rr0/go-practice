package zoo

import (
	"fmt"
	"gocourse02/utils"
)

type Zookeeper struct {
	name           string  
	holdingAnimals []Multipler
}

func NewZookeeper(name string, animals []Multipler) *Zookeeper {
	return &Zookeeper{
		name:           name,
		holdingAnimals: animals,
	}
}
func (z *Zookeeper) Name() string {
	return z.name
}
func (z *Zookeeper) HoldingAnimals() []Multipler {
	return z.holdingAnimals
}

func (z *Zookeeper) CatchAnimal(animal Multipler) {
	if utils.RandomBool() {
		z.holdingAnimals = append(z.holdingAnimals, animal)
		z.showInteraction(animal, Catch)
	} else {
		z.showInteraction(animal, Miss)
	}
}

func (z *Zookeeper) showInteraction(animal Multipler, InteractionType InteractionType) {
	fmt.Printf("Zookeeper %s %s Animal\n", z.Name(), InteractionType.String())
}
