package zoo

type InteractionType int

const (
	Catch InteractionType = iota
	Miss
	LetGo
)

func (s InteractionType) String() string {
	switch s {
	case Catch:
		return "Catch"
	case Miss:
		return "Miss"
	case LetGo:
		return "LetGo"
	}
	return "unknown"
}
