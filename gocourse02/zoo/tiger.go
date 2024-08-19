package zoo

import (
	"fmt"
	"reflect"
)

type Tiger struct {
	*Animal
	Age int
}

func NewTiger(animal *Animal, age int) *Tiger {
	return &Tiger{
		Animal: animal,
		Age:    age,
	}
}

func (t *Tiger) Reproduce(otherTiger *Tiger) *Tiger {
	age := 0
	return NewTiger(t.Animal.Reproduce(otherTiger.Animal), age)
}

func (t *Tiger) Show() {
	fmt.Println(reflect.TypeOf(t))
	t.Animal.ShowAnimal()
}
