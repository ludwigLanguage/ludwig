package vm

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
	"math"
)

func (v *VM) evalBuildList(location int) int {
	listLength := int(bytecode.ReadUint16(v.currentFrame().Instructions()[location+1:]))
	location += 2

	listEntries := []values.Value{}
	for i := 1; i <= listLength; i++ {
		/* Due to the "first in, last out" nature of the stack,
		 * we must pre-pend the new value to the list in order to
		 * keep the list in the order the user intended it to be
		 */
		listEntries = append([]values.Value{v.pop()}, listEntries...)
	}

	list := values.List{listEntries}
	v.push(list)
	return location
}

func (v *VM) evalSlice(location int) int {
	end := v.pop()
	start := v.pop()
	list := v.pop()

	startIndex, endIndex := v.checkedSliceIndices(start, end)

	var result values.Value
	switch list.Type() {
	case values.LIST:
		result = v.sliceList(list, startIndex, endIndex)
	case values.STR:
		result = v.sliceStr(list, startIndex, endIndex)
	default:
		v.raiseError("Type", "Can only slice list or string")
	}

	v.push(result)
	return location
}

func (v *VM) checkedSliceIndices(start, end values.Value) (int, int) {
	if start.Type() != values.NUM || end.Type() != values.NUM {
		v.raiseError("Type", "Can only slice from number to number")
	}

	startIndex := start.(values.Number).Value
	endIndex := end.(values.Number).Value

	if math.Floor(startIndex) != startIndex || math.Floor(endIndex) != endIndex {
		v.raiseError("Index", "Start and end indices of slices must be whole numbers")
	}

	if startIndex > endIndex {
		v.raiseError("Index", "Starting index must be smaller than ending index")
	}

	if startIndex < 0 {
		v.raiseError("Index", "Index out of range")
	}

	return int(startIndex), int(endIndex)
}

func (v *VM) sliceList(list values.Value, start, end int) values.Value {
	listEntries := list.(values.List).Values

	if end > len(listEntries)-1 {
		v.raiseError("Index", "Index out of range")
	}

	vals := listEntries[start:end]
	return values.List{vals}
}

func (v *VM) sliceStr(str values.Value, start, end int) values.Value {
	strVal := str.(values.String).Value

	if end > len(strVal)-1 {
		v.raiseError("Index", "Index out of range")
	}

	vals := strVal[start:end]
	return values.String{vals}
}

func (v *VM) evalIndex(location int) int {
	index := v.pop()
	src := v.pop()

	if index.Type() != values.NUM {
		v.raiseError("Type", "Indices may only take numbers")
	}

	indexVal := index.(values.Number).Value
	if indexVal < 0 {
		v.raiseError("Index", "Index out of range")
	}

	if math.Floor(indexVal) != indexVal {
		v.raiseError("Index", "May only index on whole numbers")
	}
	startVal := int(indexVal)

	var result values.Value
	switch src.Type() {
	case values.STR:
		result = v.sliceStr(src, startVal, startVal+1)
	case values.LIST:
		slice := v.sliceList(src, startVal, startVal+1)
		result = slice.(values.List).Values[0]
	default:
		v.raiseError("Type", "Cannot take index on this type")
	}
	v.push(result)

	return location
}
