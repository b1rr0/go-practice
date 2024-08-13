package main

import (
	"fmt"
	zoo "gocourse02/zoo"
)

func main() {
	lion := zoo.NewLion(zoo.NewAnimal("Mykola", zoo.Man), "green")
	lion2 := zoo.NewLion(zoo.NewAnimal("Natalia", zoo.Woman), "white")
	sonOfLions := lion.Multiply(*lion2)
	sonOfLions.ShowLion()


	zooKeper := zoo.NewZookeeper("Mykyta", []zoo.Multipler{})
    zooKeper.CatchAnimal(lion.Animal)
	zooKeper.CatchAnimal(lion2.Animal)
	zooKeper.CatchAnimal(sonOfLions.Animal)

	zoo.NewAnimal("Krakozyabra", zoo.Man); 
    
	for _, animal := range zooKeper.HoldingAnimals() {
		fmt.Printf("animal.Sex().String(): %v\n", animal.Sex().String())
		fmt.Printf("animal.Name(): %v\n", animal.Name())
	}
}

