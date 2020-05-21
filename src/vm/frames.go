package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

type Frame struct {
	fn               values.Function
	localInsLocation int
}

func NewFrame(fn values.Function, outerFrame *Frame) *Frame {

	f := &Frame{}
	f.fn = fn
	f.localInsLocation = 0
	return f
}

func (f *Frame) Instructions() bytecode.Instructions {
	return f.fn.Instructions
}

///////////////////////////////////////////////////////////////

func (v *VM) currentFrame() *Frame {
	return v.frames[v.framePointer-1]
}

func (v *VM) pushFrame(f *Frame) {
	v.frames[v.framePointer] = f
	v.framePointer++
}

func (v *VM) popFrame() *Frame {
	f := v.currentFrame()
	v.framePointer--
	return f
}
