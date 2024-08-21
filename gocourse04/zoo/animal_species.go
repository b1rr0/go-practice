package zoo

type AnimalSpecies int

const (
	Snake AnimalSpecies = iota
	Dog
	Cat
	Zibraa
	Deer
	Elephant
	Flamingo
	Giraffe
	Horse
	Kangaroo
	Monkey
)

func (a AnimalSpecies) String() string {
	return [...]string{"Snake", "Dog", "Cat", "Zibraa", "Deer", "Elephant", "Flamingo", "Giraffe", "Horse", "Kangaroo", "Monkey"}[a]
}

func AllAnimalSpeciesValues() []AnimalSpecies {
	return []AnimalSpecies{Snake, Dog, Cat, Zibraa, Deer, Elephant, Flamingo, Giraffe, Horse, Kangaroo, Monkey}
}
