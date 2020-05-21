package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

func (v *VM) evalCall(location int) int {
	calledVal := v.pop()
	lenOfCallArgs := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location += 2

	switch calledVal.Type() {
	case values.FUNC:
		v.callFn(calledVal, int(lenOfCallArgs)) //This function will push the result to the stack
	default:
		v.raiseError("Call", "Cannot call on this type "+calledVal.Stringify())
	}

	return location
}

func (v *VM) callFn(calledVal values.Value, lenOfCallArgs int) {
	fn := calledVal.(values.Function)

	if lenOfCallArgs != fn.NumOfArgs {
		v.raiseError("Argument", "invalid number of arguments")
	}

	fnFrame := NewFrame(fn, v.currentFrame())
	v.pushFrame(fnFrame)
	v.Run()
	v.popFrame()
}
