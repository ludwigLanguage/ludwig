package bytecode

import (
	"fmt"
)

type Definition struct {
	Name     string
	OpWidths []int
}

var definitions = map[OpCode]*Definition{
	LOADCONST: {"LoadConst", []int{2}},
	ADD:       {"add", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("OpCode %v is undefined", op)
	}

	return def, nil
}
