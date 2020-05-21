package vm

import (
	"ludwig/src/bytecode"
)

func (v *VM) evalSaveVal(location int) int {
	valIndex := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location += 2

	val := v.pop()
	v.variables[valIndex] = val
	v.push(val)

	return location
}

func (v *VM) evalGetVal(location int) int {
	valIndex := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location += 2
	val := v.variables[valIndex]
	v.push(val)
	return location
}
