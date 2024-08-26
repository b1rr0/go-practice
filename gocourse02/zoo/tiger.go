package zoo

import (
	"fmt"
	"reflect"
)

type Tiger struct {
	*Animal
	Age int
}

const (
	description = "Tiger something cool"
)

func NewTiger(name string, sex Sex, age int) *Tiger {
	return &Tiger{
		Animal: NewAnimal(name, sex),
		Age:    age,
	}
}

func (t *Tiger) Reproduce(tiger *Tiger) *Tiger {
	age := 0
	child := t.Animal.Reproduce(tiger.Animal)
	if child == nil {
		return nil
	}
	return NewTiger(child.name, child.sex, age)
}

func (t *Tiger) Show() {
	fmt.Println(reflect.TypeOf(t))
	fmt.Print(description)
	t.Animal.Show()
}
