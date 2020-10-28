package state

import "fmt"

type State struct {
	PIR1 bool
	PIR2 bool
	PIR3 bool
	PIR4 bool
}

func (s State) Serialize() []byte {
	return []byte(fmt.Sprintf("%b%b%b%b",
		boolToInt(s.PIR1),
		boolToInt(s.PIR2),
		boolToInt(s.PIR3),
		boolToInt(s.PIR4),
	))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
