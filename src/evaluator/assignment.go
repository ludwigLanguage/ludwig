package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"fmt"
)

func evalAssignment(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	leftType := fmt.Sprintf("%T", n.Left)


	val := EvalExpr(n.Right, consts)
	
	if leftType == "*ast.Identifier" {
		id := n.Left.(*ast.Identifier)
		if id.Value == "self" || id.Value == "recurse" {
			message.RaiseError("Assignment", "Cannot assign a value to the identifier 'self' or 'recurse'", n.GetTok())
		}

		consts.SetVal(id.Value, val)
	} else if leftType == "*ast.Index" {
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
        if !(length - 1 >= indexInt) {
        	message.RaiseError("Index", "Index out of range", n.GetTok())
        }

		if sourceList.Type() == values.STR {
			str := sourceList.(*values.String).Value
			str = str[0:indexInt] + val.Stringify() + str[indexInt+1:length]
			val = &values.String {str, n.GetTok()}
			
			if fmt.Sprintf("%T", index.Src) == "*ast.Identifier" {
				consts.SetVal(index.Src.(*ast.Identifier).Value, val)
			}
			
		} else if sourceList.Type() == values.LIST {
			sourceList.(*values.List).Values[indexInt] = val
		}

	} else if leftType == "*ast.InfixExpr" {
		if n.Left.(*ast.InfixExpr).Op != "." {
			message.RaiseError("Syntax", "Cannot assign to an expression", n.GetTok())
		}

		obj := EvalExpr(n.Left.(*ast.InfixExpr).Left, consts)
		if obj.Type() != values.OBJ {
			message.RaiseError("Syntax", "Cannot reassign value in non-object", n.GetTok())
		}
		id := n.Left.(*ast.InfixExpr).Right
		idType := fmt.Sprintf("%T", id)
		if idType != "*ast.Identifier" {
			message.RaiseError("Syntax", "Cannot assign to non-identifer", n.GetTok())
		}

		obj.(*values.Object).Consts.SetVal(id.(*ast.Identifier).Value, val)

	} else {
		message.RaiseError("Syntax", "Must have an identifier on left side of '='", n.GetTok())
	}


	return val
}

