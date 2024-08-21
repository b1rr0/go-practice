package main

import (
	"fmt"

	"gocourse02/zoo"
)

func main() {
	lion := zoo.NewLion("Mykola", zoo.Male, "green")
	lioness := zoo.NewLion("Natalia", zoo.Female, "white")
	cub := lion.Reproduce(*lioness)
	cub.Show()

	zookeeper := zoo.NewZookeeper("Mykyta", []zoo.Reproducer{})
	zookeeper.CatchAnimal(lion.Animal)
	zookeeper.CatchAnimal(lioness.Animal)
	zookeeper.CatchAnimal(cub.Animal)
	krakozyabra := zoo.NewAnimal("Krakozyabra", zoo.Male)
	krakozyabra.Show()

	for _, animal := range zookeeper.HoldingAnimals() {
		fmt.Printf("animal.Sex().String(): %v\n", animal.Sex().String())
		fmt.Printf("animal.Name(): %v\n", animal.Name())
	}
}
