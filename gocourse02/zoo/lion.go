package zoo

import (
	"fmt"
	"reflect"
)

type Lion struct {
	*Animal
	color string
}

func (l *Lion) Color() string {
	return l.color
}

func NewLion(animal *Animal, color string) *Lion {
	return &Lion{
		Animal: animal,
		color:  color,
	}
}

func (l *Lion) Reproduce(lion Lion) *Lion {
	color := l.Color() + "&" + lion.Color()
	return NewLion(l.Animal.Reproduce(lion.Animal), color)
}

func (l *Lion) Show() {
	fmt.Println(reflect.TypeOf(l))
	l.Animal.ShowAnimal()
}
