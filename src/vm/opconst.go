package vm

import (
	"ludwig/src/bytecode"
)

func (v *VM) evalOpConst(location int) int {
	constIndex := bytecode.ReadUint16(v.instructions[location+1:])
	location += 2

	v.push(v.pool[constIndex])

	return location
}
