package bytecode

import (
	"encoding/binary"
)

func MakeInstruction(op OpCode, args ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	lengthOfInstructions := 1
	for _, w := range def.OpWidths {
		lengthOfInstructions += w
	}

	instruction := make([]byte, lengthOfInstructions)
	instruction[0] = byte(op)
	offset := 1
	for argNumber, arg := range args {
		argWidth := def.OpWidths[argNumber]
		if argWidth == 2 {
			binary.BigEndian.PutUint16(instruction[offset:], uint16(arg))
		}
		offset += argWidth
	}

	return instruction
}
