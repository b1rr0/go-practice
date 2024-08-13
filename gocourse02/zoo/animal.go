package zoo

import (
	"fmt"
	"gocourse02/utils"
)

type Animal struct {
	name string
	sex  Sex
}

type Multipler interface {
	Sex() Sex
	Name() string
	Multiply(Multipler) *Animal
}

func (animal *Animal) Name() string {
	return animal.name
}
func (animal *Animal) Sex() Sex {
	return animal.sex
}

func NewAnimal(name string, sex Sex) *Animal {
	return &Animal{
		name: name,
		sex:  sex,
	}
}

func (animal *Animal) Multiply(otherAnimal Multipler) *Animal {
	if otherAnimal == nil {
		fmt.Printf("Can't multiply %s and nil\n", animal.Name())
		return nil
	}
	if animal.Sex() == otherAnimal.Sex() {
		fmt.Printf("Can't multiply %s and %s\n because they have the same sex", animal.Name(), otherAnimal.Name())
		return nil
	}

	newAnimalName := fmt.Sprintf("%s&%s", animal.Name(), otherAnimal.Name())
	return NewAnimal(newAnimalName, Sex(utils.RandomOneZero()))
}

func (a *Animal) ShowAnimal() {
	fmt.Println(a.Name() + " " + a.Sex().String())
}
