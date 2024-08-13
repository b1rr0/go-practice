package zoo

import (
	"fmt"
	"gocourse02/utils"
)

type Snake struct {
	*Animal
	Length int `json:"length" bson:"length"`
}

func NewSnake(animal *Animal, length int) *Snake {
	return &Snake{
		Animal: animal,
		Length: length,
	}
}

func (t *Snake) Multiply(otherSnake *Snake) *Snake {
	length := 1
	return NewSnake(t.Animal.Multiply(otherSnake.Animal), length)
}

func (t *Snake) ShowSnake() {
	fmt.Println(utils.GetType(t))
	t.Animal.ShowAnimal()
}