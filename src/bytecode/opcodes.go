package bytecode

type OpCode byte

const (
	LOADCONST OpCode = iota
	POP

	ADD
	SUB
	MULT
	DIV
	POW
)
