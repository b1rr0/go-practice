package zoo

import (
	"fmt"

	"gocourse02/utils"
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

func (l *Lion) Multiply(otherLion Lion) *Lion {
	color := l.Color() + "&" + otherLion.Color()
	return NewLion(l.Animal.Multiply(otherLion.Animal), color)
}

func (l *Lion) ShowLion() {
	fmt.Println(utils.GetType(l))
	l.Animal.ShowAnimal()
}
