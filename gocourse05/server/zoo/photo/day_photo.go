package photo

import (
	"fmt"
	"gocourse05/server/zoo"
	"math/rand"
)

type NightPhoto struct {
	cameraInfo string
	ID         string
	X          int
	Y          int
	Action     zoo.Action
	Animal     zoo.AnimalType
}

func NewNightPhoto() Photo {
	return &NightPhoto{}
}

func RandomNightPhoto(cameraInfo string) NightPhoto {
	return NightPhoto{
		cameraInfo: cameraInfo,
		ID:         fmt.Sprintf("%d", rand.Int()),
		X:          rand.Intn(100) + 1,
		Y:          rand.Intn(100) + 1,
		Action:     zoo.AllActionEnamValues()[rand.Intn(len(zoo.AllActionEnamValues()))],
		Animal:     zoo.AllAnimalTypeEnamValues()[rand.Intn(len(zoo.AllAnimalTypeEnamValues()))],
	}
}

func (c NightPhoto) PhotoVeiw() string {
	return fmt.Sprintf("Photo of %s with %s power:%d chargin:%d at (%s, %d)", c.Animal, c.Action, c.X, c.Y, c.cameraInfo, c.ID)
}
