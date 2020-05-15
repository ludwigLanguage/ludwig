package vm

import (
	"ludwig/src/values"
)

func (v *VM) add(location int) int {
	right := v.pop()
	left := v.pop()
	rVal := right.(*values.Number).Value
	lVal := left.(*values.Number).Value

	result := &values.Number{lVal + rVal, left.GetTok()}
	v.push(result)

	return location //the location does not change during additon
}
