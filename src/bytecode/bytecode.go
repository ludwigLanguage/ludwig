package bytecode

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (i Instructions) String() string {
	var out bytes.Buffer

	iter := 0
	for iter < len(i) {
		def, err := Lookup(i[iter])
		if err != nil {
			fmt.Fprintf(&out, "Error: %s\n", err)
			continue
		}

		args, read := ReadArgs(def, i[iter+1:])
		fmt.Fprintf(&out, "%04d %s\n", iter, i.fmtInstruction(def, args))
		iter += 1 + read
	}

	return out.String()
}

func (i Instructions) fmtInstruction(def *Definition, args []int) string {
	argCount := len(def.OpWidths)

	if len(args) != argCount {
		return fmt.Sprintf("Error: arg len %d does not matched defined len %d\n", argCount, len(args))
	}

	switch argCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, args[0])
	}

	return fmt.Sprintf("Error: Unhandled argCount for %s\n", def.Name)
}

func ReadArgs(def *Definition, ins Instructions) ([]int, int) {
	args := make([]int, len(def.OpWidths))
	offset := 0

	for i, width := range def.OpWidths {
		switch width {
		case 2:
			args[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}

	return args, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
