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

	} else if src.Type() != values.LIST {
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
