package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalList(n *ast.List, consts *values.SymTab) values.Value {
	l := &values.List{}

	for _, i := range n.Entries {
		l.Values = append(l.Values, EvalExpr(i, consts))
	}

	l.Tok = n.Tok
	return l
}

func evalIndex(n *ast.Index, consts *values.SymTab) values.Value {
	src := EvalExpr(n.Src, consts)
	indexVal := EvalExpr(n.Index, consts)

	if indexVal.Type() != values.NUM {
		message.RaiseError("Type", "Must have a number for indexing", indexVal.GetTok())
	}
	index := int(indexVal.(*values.Number).Value)

	if src.Type() == values.STR {
		return sliceStr(src.(*values.String), index)

	}

	if src.Type() != values.LIST {
		message.RaiseError("Type", "Must have list for indexing", src.GetTok())
	}

	lst := src.(*values.List)

	if index < 0 || index > len(lst.Values)-1 {
		message.RaiseError("Index", "Index out of range", indexVal.GetTok())
	}

	return lst.Values[index]
}

func sliceStr(v *values.String, index int) values.Value {
	if index < 0 || index > len(v.Value)-1 {
		message.RaiseError("Index", "Index out of range", v.GetTok())
	}

	return &values.String{string(v.Value[index]), v.GetTok()}
}

func evalSlice(n *ast.Slice, consts *values.SymTab) values.Value {
	source := EvalExpr(n.Src, consts)
	startVal := EvalExpr(n.Start, consts)

	var endVal values.Value
	if n.End == nil {
		endVal = values.Length([]values.Value{source}, n.GetTok())
		endVal.(*values.Number).Value--
	} else {
		endVal = EvalExpr(n.End, consts)
	}

	if startVal.Type() != values.NUM && endVal.Type() != values.NUM {
		message.RaiseError("Type", "The start value and end value of a slice must both be numbers", n.GetTok())
	}
	start := int(startVal.(*values.Number).Value)
	end := int(endVal.(*values.Number).Value)

	var rtrnVal values.Value

	if source.Type() == values.STR {
		str := source.(*values.String).Value
		if end > len(str)-1 || start < 0 {
			message.RaiseError("Index", "Index out of range", n.GetTok())
		}

		val := str[start:end]
		rtrnVal = &values.String{val, n.GetTok()}

	} else if source.Type() == values.LIST {
		lst := source.(*values.List).Values
		if end > len(lst)-1 || start < 0 {
			message.RaiseError("Index", "Index out of range", n.GetTok())
		}

		val := lst[start:end]
		rtrnVal = &values.List{val, n.GetTok()}

	} else {
		message.RaiseError("Type", "Must either have a string or list for slicing", n.Tok)
	}

	return rtrnVal
}
