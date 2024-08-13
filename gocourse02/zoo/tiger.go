package zoo

import (
	"fmt"
	"gocourse02/utils"
)

type Tiger struct {
	*Animal
	Age int `json:"age" bson:"age"`
}

func NewTiger(animal *Animal, age int) *Tiger {
	return &Tiger{
		Animal: animal,
		Age:    age,
	}
}

func (t *Tiger) Multiply(otherTiger *Tiger) *Tiger {
	age := 0
	return NewTiger(t.Animal.Multiply(otherTiger.Animal), age)
}

func (t *Tiger) ShowTiger() {
	fmt.Println(utils.GetType(t))
	t.Animal.ShowAnimal()
}