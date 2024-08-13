package zoo

type Sex int

const (
	Man Sex = iota
	Woman
)

func (s Sex) String() string {
	switch s {
	case Man:
		return "Man"
	case Woman:
		return "Woman"
	}
	return "unknown"
}
