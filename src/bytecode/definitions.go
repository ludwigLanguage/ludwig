package bytecode

import (
	"fmt"
)

type Definition struct {
	Name     string
	OpWidths []int
}

var definitions = map[OpCode]*Definition{
	LOADCONST: {"load constant", []int{2}},
	POP:       {"pop stack", []int{}},
	ADD:       {"add", []int{}},
	SUB:       {"subtract", []int{}},
	MULT:      {"multiply", []int{}},
	DIV:       {"divide", []int{}},
	POW:       {"power", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("OpCode %v is undefined", op)
	}

	return def, nil
}
