package zoo

import (
	"fmt"
	"reflect"
)

type Snake struct {
	*Animal
	Length int
}

func NewSnake(animal *Animal, length int) *Snake {
	return &Snake{
		Animal: animal,
		Length: length,
	}
}

func (t *Snake) Reproduce(otherSnake *Snake) *Snake {
	length := 1
	return NewSnake(t.Animal.Reproduce(otherSnake.Animal), length)
}

func (t *Snake) Show() {
	fmt.Println(reflect.TypeOf(t))
	t.Animal.ShowAnimal()
}
