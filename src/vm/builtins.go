package vm

import (
	"fmt"
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

type builtinFn func(v []values.Value) values.Value

var builtinsMap = map[int]builtinFn{
	bytecode.PRINT:   print,
	bytecode.PRINTLN: println,
}

func (v *VM) evalBuiltin(location int) int {
	id := bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:])
	location += 2

	fn := builtinsMap[int(id)]
	if fn == nil {
		v.raiseError("Implementation", "This builtin is not implemented")
	}

	vals := make([]values.Value, v.stackPointer)
	for v.stackPointer > 0 {
		vals[v.stackPointer] = v.pop()
	}

	v.push(fn(vals))
	return location
}

func print(v []values.Value) values.Value {
	var out string

	vLen := len(v)
	for j, i := range v {
		out += i.Stringify()

		if j != vLen-1 {
			out += " "
		}
	}

	fmt.Print(out)
	return values.String{out}
}

func println(v []values.Value) values.Value {
	newline := values.String{"\n"}
	return print(append(v, newline))
}
