package zoo

type AnimalType int

const (
	Snake AnimalType = iota
	Lion
	Tiger
	Elephant
	Giraffe
	Zebra
	Deer
	Dog
	Cat
	Monkey
)

const DefaultAnimalType = Snake

func (a AnimalType) String() string {
	return [...]string{
		"Snake",
		"Lion",
		"Tiger",
		"Elephant",
		"Giraffe",
		"Zebra",
		"Deer",
		"Dog",
		"Cat",
		"Monkey",
	}[a]
}

func AllAnimalTypeEnamValues() []AnimalType {
	return []AnimalType{
		Snake,
		Lion,
		Tiger,
		Elephant,
		Giraffe,
		Zebra,
		Deer,
		Dog,
		Cat,
		Monkey,
	}
}
