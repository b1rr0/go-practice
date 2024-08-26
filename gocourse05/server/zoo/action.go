package zoo

type Action int

const (
	Jamp Action = iota
	Sleep
	Run
	Ester
	Play
	Eat
	Drink
	Swim
)

const DefaultAction = Jamp

func (a Action) String() string {
	return [...]string{
		"Jamp",
		"Sleep",
		"Run",
		"Ester",
		"Play",
		"Eat",
		"Drink",
		"Swim",
		"Clap",
		"Dance",
		"Laugh",
		"Cry",
		"Shout",
		"Whistle",
	}[a]
}

func AllActionEnamValues() []Action {
	return []Action{
		Jamp,
		Sleep,
		Run,
		Ester,
		Play,
		Eat,
		Drink,
		Swim,
	}
}
