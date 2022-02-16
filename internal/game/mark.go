package game

// Mark represents a game marking (X or O).
type Mark bool

// ToString returns a string representation of a Mark.
func (m *Mark) ToString() string {
	if m == nil {
		return "-"
	}
	if *m {
		return "X"
	}
	return "O"
}

// Because Go doesn't support enums or unions in general, this is an OK
// workaround.
const (
	X = Mark(true)
	O = Mark(false)
)
