package zoo

import (
	"fmt"
)

type Animal struct {
	name string
	sex  Sex
}

type Reproducer interface {
	Sex() Sex
	Name() string
	Reproduce(Reproducer) *Animal
}

func (a *Animal) Name() string {
	return a.name
}
func (a *Animal) Sex() Sex {
	return a.sex
}

func NewAnimal(name string, sex Sex) *Animal {
	return &Animal{
		name: name,
		sex:  sex,
	}
}

func (a *Animal) Reproduce(animal Reproducer) *Animal {
	if animal == nil {
		fmt.Printf("Can't multiply %s and nil\n", a.Name())
		return nil
	}
	if a.Sex() == animal.Sex() {
		fmt.Printf("Can't multiply %s and %s\n because they have the same sex", a.Name(), animal.Name())
		return nil
	}

	newAnimalName := fmt.Sprintf("%s&%s", a.Name(), animal.Name())
	return NewAnimal(newAnimalName, Sex(RandomOneZero()))
}

func (a *Animal) ShowAnimal() {
	fmt.Println(a.Name() + " " + a.Sex().String())
}
