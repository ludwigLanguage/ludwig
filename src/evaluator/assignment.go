package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
)

func evalAssignment(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	leftType := n.Left.Type()

	if leftType == ast.IDENT {
		return evalIdentAssignment(n, consts)

	} else if leftType == ast.INDEX {
		return evalIndexAssignment(n, consts)

	} else if leftType == ast.INFIX {
		return evalDotAssignment(n, consts)

	} else {
		message.RaiseError("Syntax", "Must have an identifier on left side of '='", n.GetTok())
	}

	return nil //safe because message.RaiseError exits program
}

func evalIdentAssignment(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	val := EvalExpr(n.Right, consts)
	id := n.Left.(*ast.Identifier)
	if id.Value == "self" || id.Value == "recurse" {
		message.RaiseError("Assignment", "Cannot assign a value to the identifier 'self' or 'recurse'", n.GetTok())
	}

	consts.SetVal(id.Value, val)

	self := consts.GetVal("self")
	if self != nil {
		self.(*values.Object).Consts.SetVal(id.Value, val)
	}
	return val
}

func evalIndexAssignment(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	val := EvalExpr(n.Right, consts)
	index := n.Left.(*ast.Index)
	sourceList := EvalExpr(index.Src, consts)
	indexVal := EvalExpr(index.Index, consts)

	length := 0
	if sourceList.Type() == values.STR {
		length = len(sourceList.(*values.String).Value)
	} else if sourceList.Type() == values.LIST {
		length = len(sourceList.(*values.List).Values)
	} else {
		message.RaiseError("Type", "Can only edit an index of a string or a list", n.GetTok())
	}

	indexInt := int(indexVal.(*values.Number).Value)
	if !(length-1 >= indexInt) {
		message.RaiseError("Index", "Index out of range", n.GetTok())
	}

	if sourceList.Type() == values.STR {
		str := sourceList.(*values.String).Value
		str = str[0:indexInt] + val.Stringify() + str[indexInt+1:length]
		val = &values.String{str, n.GetTok()}

		if index.Src.Type() == ast.IDENT {
			consts.SetVal(index.Src.(*ast.Identifier).Value, val)
		}

	} else if sourceList.Type() == values.LIST {
		sourceList.(*values.List).Values[indexInt] = val
	}

	return val
}

func evalDotAssignment(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	val := EvalExpr(n.Right, consts)

	if n.Left.(*ast.InfixExpr).Op != "." {
		message.RaiseError("Syntax", "Cannot assign to an expression", n.GetTok())
	}

	obj := EvalExpr(n.Left.(*ast.InfixExpr).Left, consts)
	if obj.Type() != values.OBJ {
		message.RaiseError("Syntax", "Cannot reassign value in non-object", n.GetTok())
	}
	id := n.Left.(*ast.InfixExpr).Right
	if id.Type() != ast.IDENT {
		message.RaiseError("Syntax", "Cannot assign to non-identifer", n.GetTok())
	}

	obj.(*values.Object).Consts.SetVal(id.(*ast.Identifier).Value, val)
	return val
}
