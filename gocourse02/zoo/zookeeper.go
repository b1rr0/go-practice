package zoo

import (
	"fmt"
)

type Zookeeper struct {
	name           string
	holdingAnimals []Reproducer
}

func NewZookeeper(name string, animals []Reproducer) *Zookeeper {
	return &Zookeeper{
		name:           name,
		holdingAnimals: animals,
	}
}

func (z *Zookeeper) Name() string {
	return z.name
}

func (z *Zookeeper) HoldingAnimals() []Reproducer {
	return z.holdingAnimals
}

func (z *Zookeeper) CatchAnimal(animal Reproducer) {
	if RandomBool() {
		z.holdingAnimals = append(z.holdingAnimals, animal)
		z.showInteraction(animal, Catch)
	} else {
		z.showInteraction(animal, Miss)
	}
}

func (z *Zookeeper) showInteraction(animal Reproducer, InteractionType InteractionType) {
	fmt.Printf("Zookeeper %s %s Animal\n", z.Name(), InteractionType.String())
}
