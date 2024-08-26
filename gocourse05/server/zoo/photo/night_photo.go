package photo

import (
	"fmt"
	"gocourse05/server/zoo"
	"math/rand"
)

type DayPhoto struct {
	cameraInfo string
	ID         int
	Animal     zoo.AnimalType
	Action     zoo.Action
	Power      int
	Chargin    int
}

func RandomDayPhoto(cameraInfo string) DayPhoto {
	return DayPhoto{
		cameraInfo: cameraInfo,
		ID:         rand.Int(),
		Animal:     zoo.AllAnimalTypeEnamValues()[rand.Intn(len(zoo.AllAnimalTypeEnamValues()))],
		Action:     zoo.AllActionEnamValues()[rand.Intn(len(zoo.AllActionEnamValues()))],
		Power:      rand.Intn(100) + 1,
		Chargin:    rand.Intn(100) + 1,
	}
}

func (c DayPhoto) PhotoVeiw() string {
	return fmt.Sprintf("Photo of %s with %s power:%d chargin:%d at (%s, %d)", c.Animal, c.Action, c.Power, c.Chargin, c.cameraInfo, c.ID)
}
