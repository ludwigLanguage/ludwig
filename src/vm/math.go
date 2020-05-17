package vm

import (
	"ludwig/src/values"
	"math"
)

func (v *VM) getType(left, right values.Value) byte {
	if left.Type() == values.NIL || right.Type() == values.NIL {
		return values.NIL
	} else if left.Type() != right.Type() {
		v.raiseError("Type", "Expected similar types on both sides of the operator")
	}

	return left.Type()
}

func (v *VM) evalAdd(location int) int {
	right := v.pop()
	left := v.pop()
	var result values.Value

	switch v.getType(left, right) {
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
	switch v.getType(left, right) {
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

func (v *VM) evalMultiply(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = multiplyNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '*' to these types")
	}

	v.push(result)
	return location
}

func multiplyNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Number{lVal * rVal}
}

func (v *VM) evalDivide(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = divideNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '/' to these types")
	}

	v.push(result)
	return location
}

func divideNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Number{lVal / rVal}
}

func (v *VM) evalPower(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = powerNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '^' to these types")
	}

	v.push(result)
	return location
}

func powerNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Number{math.Pow(lVal, rVal)}
}

func (v *VM) evalNegative(location int) int {
	right := v.pop()

	var result values.Value
	switch right.Type() {
	case values.NUM:
		result = negativeNumbers(right)
	default:
		v.raiseError("Operator", "Cannot apply '^' to these types")
	}

	v.push(result)
	return location
}

func negativeNumbers(right values.Value) values.Value {
	rVal := right.(values.Number).Value

	return values.Number{-rVal}
}

////////////////////////////////////////////////////////////////
//
// Boolean Logic ->->->->->
//
///////////////////////////////////////////////////////////////

func (v *VM) evalEqualTo(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = equalToNumbers(left, right)
	case values.BOOL:
		result = equalToBooleans(left, right)
	case values.NIL:
		result = equalToNil(left, right)
	default:
		v.raiseError("Operator", "Cannot apply '==' these types")
	}

	v.push(result)
	return location
}

func equalToNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Boolean{lVal == rVal}
}

func equalToBooleans(left, right values.Value) values.Value {
	lVal := left.(values.Boolean).Value
	rVal := right.(values.Boolean).Value

	return values.Boolean{lVal == rVal}
}

func equalToNil(left, right values.Value) values.Value {
	return values.Boolean{left.Type() == right.Type()}
}

func (v *VM) evalNotEqual(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value

	switch v.getType(left, right) {
	case values.NUM:
		result = notEqualNumbers(left, right)
	case values.BOOL:
		result = notEqualBooleans(left, right)
	case values.NIL:
		result = notEqualNil(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '!='")
	}

	v.push(result)
	return location
}

func notEqualNumbers(left, right values.Value) values.Value {
	lVal, rVal := left.(values.Number).Value, right.(values.Number).Value
	return values.Boolean{lVal != rVal}
}

func notEqualBooleans(left, right values.Value) values.Value {
	lVal, rVal := left.(values.Boolean).Value, right.(values.Boolean).Value
	return values.Boolean{lVal != rVal}
}

func notEqualNil(left, right values.Value) values.Value {
	return values.Boolean{left.Type() != right.Type()}
}

func (v *VM) evalLessThan(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = lessThanNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '<'")
	}

	v.push(result)
	return location
}

func lessThanNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Boolean{lVal < rVal}
}

func (v *VM) evalGreaterThan(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = greaterThanNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '>'")
	}

	v.push(result)
	return location
}

func greaterThanNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Boolean{lVal > rVal}
}

func (v *VM) evalLessEquals(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = lessEqualsNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '<='")
	}

	v.push(result)
	return location
}

func lessEqualsNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Boolean{lVal <= rVal}
}

func (v *VM) evalGreaterLessEquals(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.NUM:
		result = greaterEqualsNumbers(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '>='")
	}

	v.push(result)
	return location
}

func greaterEqualsNumbers(left, right values.Value) values.Value {
	lVal := left.(values.Number).Value
	rVal := right.(values.Number).Value

	return values.Boolean{lVal >= rVal}
}

func (v *VM) evalOr(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.BOOL:
		result = orBooleans(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '||'")
	}

	v.push(result)
	return location
}

func orBooleans(left, right values.Value) values.Value {
	lVal := left.(values.Boolean).Value
	rVal := right.(values.Boolean).Value

	return values.Boolean{lVal || rVal}
}

func (v *VM) evalAnd(location int) int {
	right := v.pop()
	left := v.pop()

	var result values.Value
	switch v.getType(left, right) {
	case values.BOOL:
		result = andBooleans(left, right)
	default:
		v.raiseError("Operator", "Cannot apply this type to '&&'")
	}

	v.push(result)
	return location
}

func andBooleans(left, right values.Value) values.Value {
	lVal := left.(values.Boolean).Value
	rVal := right.(values.Boolean).Value

	return values.Boolean{lVal && rVal}
}

func (v *VM) evalNot(location int) int {
	right := v.pop()

	var result values.Value
	switch right.Type() {
	case values.BOOL:
		result = notBoolean(right)
	default:
		v.raiseError("Operator", "Cannot apply '!' to this type")
	}

	v.push(result)
	return location
}

func notBoolean(right values.Value) values.Value {
	rVal := right.(values.Boolean).Value
	return values.Boolean{!rVal}
}
