package vm

import (
	"ludwig/src/values"
)

func (v *VM) evalAdd(location int) int {
	right := v.pop()
	left := v.pop()
	var result values.Value

	switch right.Type() {
	case values.NUM:
		result = addNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '+' operator to these types")
	}

	v.push(result)

	return location //the location does not change during additon
}

func addNumbers(left values.Value, right values.Value) values.Value {
	rVal := right.(values.Number).Value
	lVal := left.(values.Number).Value

	return values.Number{rVal + lVal}
}

func (v *VM) evalSubtract(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch right.Type() {
	case values.NUM:
		result = subtractNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '-' to these types")
	}

	v.push(result)
	return location
}

func subtractNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Number{lVal - rVal}
}
