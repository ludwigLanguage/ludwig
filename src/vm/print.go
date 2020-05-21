package vm

import (
	"fmt"
	"ludwig/src/values"
)

func (v *VM) evalPrint(location int) int {
	vals := []values.Value{}

	for v.stackPointer != 0 {
		vals = append(vals, v.pop())
	}

	str := ""
	for _, i := range vals {
		str = i.Stringify() + " " + str
	}
	fmt.Printf("%v\n", str)

	val := values.String{str}
	v.push(val)

	return location
}
