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

	ADD:  {"add", []int{}},
	SUB:  {"subtract", []int{}},
	MULT: {"multiply", []int{}},
	DIV:  {"divide", []int{}},
	POW:  {"power", []int{}},

	EQUALTO:       {"equal to", []int{}},
	NOTEQUAL:      {"not equal to", []int{}},
	GREATERTHAN:   {"greater than", []int{}},
	LESSTHAN:      {"less than", []int{}},
	GREATEREQUALS: {"greater than or equal to", []int{}},
	LESSEREQUALS:  {"less than or equal to", []int{}},

	OR:  {"or", []int{}},
	AND: {"and", []int{}},

	NOT:      {"not", []int{}},
	NEGATIVE: {"negative", []int{}},

	JUMP:   {"jump", []int{2}},
	JUMPNT: {"jump if not true", []int{2}},

	SAVEV: {"set value", []int{2}},
	GETV:  {"get value", []int{2}},

	BUILDLIST: {"build list", []int{2}},
	SLICE:     {"build list", []int{}},
	INDEX:     {"index list", []int{}},

	CALL:  {"Call", []int{2}},
	PRINT: {"print", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("OpCode %v is undefined", op)
	}

	return def, nil
}
