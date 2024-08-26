package zoo

import (
	"fmt"
	"reflect"
)

type Lion struct {
	*Animal
	color string
}

const (
	description = "Lion something cool"
)

func (l *Lion) Color() string {
	return l.color
}

func NewLion(name string, sex Sex, color string) *Lion {
	return &Lion{
		Animal: NewAnimal(name, sex),
		color:  color,
	}
}

func (l *Lion) Reproduce(lion Lion) *Lion {
	color := l.Color() + "&" + lion.Color()
	child := l.Animal.Reproduce(lion.Animal)
	if child == nil {
		return nil
	}
	return NewLion(child.Name(), child.sex, color)
}

func (l *Lion) Show() {
	fmt.Println(reflect.TypeOf(l))
	fmt.Print(description)
	l.Animal.Show()
}
