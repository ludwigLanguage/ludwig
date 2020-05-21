package vm

import (
	"ludwig/src/values"
)

func (v *VM) evalCall(location int) int {
	calledVal := v.pop()

	switch calledVal.Type() {
	case values.FUNC:
		v.callFn(calledVal) //This function will push the result to the stack
	default:
		v.raiseError("Call", "Cannot call on this type "+calledVal.Stringify())
	}

	return location
}

func (v *VM) callFn(calledVal values.Value) {
	fn := calledVal.(values.Function)
	fnFrame := NewFrame(fn, v.currentFrame())
	v.pushFrame(fnFrame)
	v.Run()
	v.popFrame()
}
