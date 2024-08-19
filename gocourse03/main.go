package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Sector     string
	rodentType string
	FromTo     [2]Sector

	Movement struct {
		Time time.Time
		FromTo
		rodentId int
	}

	Rodent struct {
		ID          int
		Type        rodentType
		Sector      Sector
		UpdatedTime time.Time
	}
)

func (m *Movement) Show() {
	fmt.Print(m.rodentId, " ", m.FromTo[0], "->", m.FromTo[1], " ", m.Time.Format("2006-01-02 15:04:05"), "\n")
}

func NewRodent(id int, rodentType rodentType, sector Sector) *Rodent {
	return &Rodent{
		ID:          id,
		Type:        rodentType,
		Sector:      sector,
		UpdatedTime: time.Now(),
	}
}

func (rodent *Rodent) Show() {
	fmt.Printf("ID: %d\n", rodent.ID)
	fmt.Printf("Type: %s\n", rodent.Type)
	fmt.Printf("Sector: %s\n", rodent.Sector)
	fmt.Printf("UpdatedTime: %s\n", rodent.UpdatedTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("__________________________________________\n")
}

func NewRodentInSectorA(id int, rodentType rodentType) *Rodent {
	return NewRodent(id, rodentType, "A")
}

// set of Sectors

var rodentMap = map[int]*Rodent{
	1:  NewRodentInSectorA(1, "M1"),
	2:  NewRodentInSectorA(2, "M2"),
	3:  NewRodentInSectorA(3, "M3"),
	4:  NewRodentInSectorA(4, "M3"),
	5:  NewRodentInSectorA(5, "M2"),
	6:  NewRodentInSectorA(6, "M2"),
	7:  NewRodentInSectorA(7, "M3"),
	8:  NewRodentInSectorA(8, "M3"),
	9:  NewRodentInSectorA(9, "M4"),
	10: NewRodentInSectorA(10, "M4"),
	11: NewRodentInSectorA(11, "M5"),
	12: NewRodentInSectorA(12, "M5"),
	13: NewRodentInSectorA(13, "M5"),
	14: NewRodentInSectorA(14, "M5"),
	15: NewRodentInSectorA(15, "M5"),
	16: NewRodentInSectorA(16, "M5"),
	17: NewRodentInSectorA(17, "M5"),
	18: NewRodentInSectorA(18, "M5"),
	19: NewRodentInSectorA(19, "M5"),
	20: NewRodentInSectorA(20, "M5"),
}

var sectors = []Sector{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}

var sectorsToInt = map[Sector]int{
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13,
} // just for practice

const (
	lenSectors       = 14
	countTransitions = 66
)

func generateTransitionMatrix(transitions [][]int) [][]int {
	for i := 0; i < countTransitions; i++ {
		var from, to int
		for from == to || transitions[from][to] > 0 {
			from = rand.Intn(lenSectors)
			to = rand.Intn(lenSectors)
		}
		transitions[from][to]++
		transitions[to][from]++
	}

	for _, transitions := range transitions {
		for _, transition := range transitions {
			fmt.Print(" ", transition)
		}
		fmt.Print("\n")
	}

	return transitions
}

func generateMovements(transitions [][]int) []Movement {
	movements := []Movement{}
	for i := 1; i <= 20; i++ {
		rodent := rodentMap[i]
		for j := 0; j < 10; j++ {
			from := rodent.Sector
			var to Sector
			for {
				to = sectors[rand.Intn(lenSectors)]
				if from != to && transitions[sectorsToInt[from]][sectorsToInt[to]] > 0 {
					break
				}
			}

			time := rodent.UpdatedTime.Add(time.Duration(rand.Intn(60)) * time.Minute)
			movements = append(movements, Movement{Time: time, FromTo: FromTo{from, to}, rodentId: i})

			rodent.UpdatedTime = time
			rodent.Sector = to
		}
	}
	return movements
}

func main1() { // for exec secod file
	fmt.Print(lenSectors)
	transitions := make([][]int, lenSectors)
	for i := range transitions {
		transitions[i] = make([]int, lenSectors)
	}
	fmt.Println(transitions[13][13])
	transitions = generateTransitionMatrix(transitions)
	for _, rodent := range rodentMap {
		rodent.Show()
	}

	fmt.Printf("__________________________________________")
	movents := generateMovements(transitions) // immulate random transitions

	for _, rodent := range rodentMap {
		rodent.Show()
	}
	fmt.Printf("__________________________________________")

	for _, movement := range movents {
		movement.Show()
	}
}
