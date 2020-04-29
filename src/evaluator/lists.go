package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalList(n *ast.List, consts *values.SymTab, log *message.Log) values.Value {
	l := &values.List{}

	for _, i := range n.Entries {
		l.Values = append(l.Values, EvalExpr(i, consts, log))
	}

	l.Tok = n.Tok
	return l
}

func evalIndex(n *ast.Index, consts *values.SymTab, log *message.Log) values.Value {
	src := EvalExpr(n.Src, consts, log)
	indexVal := EvalExpr(n.Index, consts, log)

	if indexVal.Type() != values.NUM {
		message.RuntimeErr("Type", "Must have a number for indexing", indexVal.GetTok(), log)
	}
	index := int(indexVal.(*values.Number).Value)

	if src.Type() == values.STR {
		return sliceStr(src.(*values.String), index, log)

	}

	if src.Type() != values.LIST {
		message.RuntimeErr("Type", "Must have list for indexing", src.GetTok(), log)
	}

	lst := src.(*values.List)

	if index < 0 || index > len(lst.Values)-1 {
		message.RuntimeErr("Index", "Index out of range", indexVal.GetTok(), log)
	}

	return lst.Values[index]
}

func sliceStr(v *values.String, index int, log *message.Log) values.Value {
	if index < 0 || index > len(v.Value)-1 {
		message.RuntimeErr("Index", "Index out of range", v.GetTok(), log)
	}

	return &values.String{string(v.Value[index]), v.GetTok()}
}

func evalSlice(n *ast.Slice, consts *values.SymTab, log *message.Log) values.Value {
	source := EvalExpr(n.Src, consts, log)
	startVal := EvalExpr(n.Start, consts, log)

	var endVal values.Value
	if n.End == nil {
		endVal = values.Length([]values.Value{source}, n.GetTok(), log)
		endVal.(*values.Number).Value--
	} else {
		endVal = EvalExpr(n.End, consts, log)
	}

	if startVal.Type() != values.NUM && endVal.Type() != values.NUM {
		message.RuntimeErr("Type", "The start value and end value of a slice must both be numbers", n.GetTok(), log)
	}
	start := int(startVal.(*values.Number).Value)
	end := int(endVal.(*values.Number).Value)

	var rtrnVal values.Value

	if source.Type() == values.STR {
		str := source.(*values.String).Value
		if end > len(str)-1 || start < 0 {
			message.RuntimeErr("Index", "Index out of range", n.GetTok(), log)
		}

		val := str[start:end]
		rtrnVal = &values.String{val, n.GetTok()}

	} else if source.Type() == values.LIST {
		lst := source.(*values.List).Values
		if end > len(lst)-1 || start < 0 {
			message.RuntimeErr("Index", "Index out of range", n.GetTok(), log)
		}

		val := lst[start:end]
		rtrnVal = &values.List{val, n.GetTok()}

	} else {
		message.RuntimeErr("Type", "Must either have a string or list for slicing", n.Tok, log)
	}

	return rtrnVal
}
