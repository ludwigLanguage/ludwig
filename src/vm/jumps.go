package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

func (v *VM) evalJump(location int) int {
	target := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location = int(target) - 1

	return location
}

func (v *VM) evalJumpIfNotTrue(location int) int {
	val := v.pop()

	if val.Type() != values.BOOL {
		v.raiseError("Type", "Expected boolean")
	}

	if val.(values.Boolean).Value {
		location += 2
	} else {
		target := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
		location = int(target) - 1
	}

	return location
}
