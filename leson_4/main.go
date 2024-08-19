/*
Беремо той самий ящик ліків із минулого заняття, але тепер дані в мапах.

На вході маємо великий ящик із ліками, на виході — три маленьких і пустий початковий.
За вимогою логіста, засовуємо три нові ящика в один великий, який вміє шукати в цих трьох маленьких.
*/

package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	Name         string
	Manufactured time.Time
}

type Box struct {
	ToDiscard map[string]Medicine
	ForSale   map[string]Medicine
	ToUse     map[string]Medicine
}

func NewBox() *Box {
	return &Box{
		ToDiscard: make(map[string]Medicine),
		ForSale:   make(map[string]Medicine),
		ToUse:     make(map[string]Medicine),
	}
}

func (b Box) FindMedicine(name string) (*Medicine, string) {
	if med, ok := b.ToDiscard[name]; ok {
		return &med, "ToDiscard"
	} else if med, ok := b.ForSale[name]; ok {
		return &med, "ForSale"
	} else if med, ok := b.ToUse[name]; ok {
		return &med, "ToUse"
	} else {
		return nil, ""
	}
}

func (b Box) AddToDiscard(med *Medicine) {
	b.ToDiscard[med.Name] = *med
}

func (b Box) AddToSale(med *Medicine) {
	b.ForSale[med.Name] = *med
}

func (b Box) AddToUse(med *Medicine) {
	b.ToUse[med.Name] = *med
}

func (b Box) FilterMedicines(medicines []Medicine) {

	for _, med := range medicines {
		if med.Manufactured.Before(time.Now()) {
			b.AddToDiscard(&med)
		} else if med.Manufactured.After(time.Now()) {
			b.AddToSale(&med)
		} else {
			b.AddToUse(&med)
		}
	}
}
func main() {
	medicines := []Medicine{
		{Name: "n1", Manufactured: time.Now().AddDate(0, 0, -1)},
		{Name: "n2", Manufactured: time.Now().AddDate(0, 0, -260)},
		{Name: "n3", Manufactured: time.Now().AddDate(0, 0, 1)},
		{Name: "n4", Manufactured: time.Now().AddDate(0, 0, 1)},
		{Name: "n5", Manufactured: time.Now().AddDate(0, 0, 1)},
		{Name: "n6", Manufactured: time.Now().AddDate(0, 0, 1)},
		{Name: "n7", Manufactured: time.Now().AddDate(0, 0, 1)},
		{Name: "n8", Manufactured: time.Now()},
		{Name: "n9", Manufactured: time.Now()},
		{Name: "n10", Manufactured: time.Now()},
		{Name: "n11", Manufactured: time.Now()},
	}

	box := NewBox()
	box.FilterMedicines(medicines)

	if med, category := box.FindMedicine("n1"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'n1' не знайдені.")
	}

	if med, category := box.FindMedicine("n10"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'n10' не знайдені.")
	}

	if med, category := box.FindMedicine("n11"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'n11' не знайдені.")
	}

	if med, category := box.FindMedicine("n5"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'n5' не знайдені.")
	}

}
