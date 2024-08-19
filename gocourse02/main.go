package main

import (
	"fmt"
	zoo "gocourse02/zoo"
)

func main() {
	lion := zoo.NewLion(zoo.NewAnimal("Mykola", zoo.Male), "green")
	lioness := zoo.NewLion(zoo.NewAnimal("Natalia", zoo.Female), "white")
	cub := lion.Reproduce(*lioness)
	cub.Show()

	zooKeeper := zoo.NewZookeeper("Mykyta", []zoo.Reproducer{})
	zooKeeper.CatchAnimal(lion.Animal)
	zooKeeper.CatchAnimal(lioness.Animal)
	zooKeeper.CatchAnimal(cub.Animal)
	krakozyabra := zoo.NewAnimal("Krakozyabra", zoo.Male)
	krakozyabra.ShowAnimal()

	for _, animal := range zooKeeper.HoldingAnimals() {
		fmt.Printf("animal.Sex().String(): %v\n", animal.Sex().String())
		fmt.Printf("animal.Name(): %v\n", animal.Name())
	}
}
