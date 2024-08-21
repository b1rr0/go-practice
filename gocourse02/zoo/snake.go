package zoo

import (
	"fmt"
	"reflect"
)

type Snake struct {
	*Animal
	Length int
}

const (
	defaultLength = 1
	description   = "Snake something cool"
)

func NewSnake(name string, sex Sex, length int) *Snake {
	return &Snake{
		Animal: NewAnimal(name, sex),
		Length: length,
	}
}

func (s *Snake) Reproduce(snake *Snake) *Snake {
	child := s.Animal.Reproduce(snake.Animal)
	if child == nil {
		return nil
	}
	return NewSnake(child.name, child.sex, defaultLength)
}

func (s *Snake) Show() {
	fmt.Println(reflect.TypeOf(s))
	fmt.Print(description)
	s.Animal.Show()
}
