package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"math"
)

func evalInfix(n *ast.InfixExpr, consts *values.SymTab) values.Value {

	if n.Op == "=" {
		return evalAssignment(n, consts)
	} else if n.Op == "." {
		return evalObjInfix(n, consts)
	}

	leftVal := EvalExpr(n.Left, consts)

	rightVal := EvalExpr(n.Right, consts)
	if leftVal.Type() == values.NIL || rightVal.Type() == values.NIL {
		isTheSame := leftVal.Type() == rightVal.Type()
		if n.Op == "==" {
			return &values.Boolean{isTheSame, n.GetTok()}
		} else if n.Op == "!=" {
			return &values.Boolean{!isTheSame, n.GetTok()}
		} else {
			message.RaiseError("Type", "Cannot evaluate an infix expression for these types", n.GetTok())
		}
	}

	switch leftVal.Type() {
	case values.NUM:
		return evalNumInfix(leftVal, rightVal, n.Op)
	case values.BOOL:
		return evalBoolInfix(leftVal, rightVal, n.Op)
	case values.STR:
		return evalStrInfix(leftVal, rightVal, n.Op)
	case values.LIST:
		return evalListInfix(leftVal, rightVal, n.Op)
	default:
		message.RaiseError("Type", "Cannot evaluate an infix expression for these types", n.GetTok())
	}
	return NIL
}

func evalNumInfix(l, r values.Value, op string) values.Value {
	if l.Type() != r.Type() {
		message.RaiseError("Type", "Invalid right side type", r.GetTok())
	}
	left := l.(*values.Number)
	right := r.(*values.Number)

	switch op {
	case "+":
		return &values.Number{left.Value + right.Value, left.GetTok()}
	case "-":
		return &values.Number{left.Value - right.Value, left.GetTok()}
	case "*":
		return &values.Number{left.Value * right.Value, left.GetTok()}
	case "/":
		return &values.Number{left.Value / right.Value, left.GetTok()}
	case "^":
		return &values.Number{math.Pow(left.Value, right.Value), left.GetTok()}
	case "==":
		return &values.Boolean{left.Value == right.Value, left.GetTok()}
	case "!=":
		return &values.Boolean{left.Value != right.Value, left.GetTok()}
	case "<":
		return &values.Boolean{left.Value < right.Value, left.GetTok()}
	case ">":
		return &values.Boolean{left.Value > right.Value, left.GetTok()}
	case ">=":
		return &values.Boolean{left.Value >= right.Value, left.GetTok()}
	case "<=":
		return &values.Boolean{left.Value <= right.Value, left.GetTok()}
	default:
		message.RaiseError("Operator", "This operator is not defined on numbers", left.GetTok())
	}

	return NIL
}

func evalBoolInfix(l, r values.Value, op string) values.Value {
	if l.Type() != r.Type() {
		message.RaiseError("Type", "Invalid type on the right side", r.GetTok())
	}
	left := l.(*values.Boolean)
	right := r.(*values.Boolean)

	switch op {
	case "==":
		return &values.Boolean{left.Value == right.Value, l.GetTok()}
	case "!=":
		return &values.Boolean{left.Value != right.Value, l.GetTok()}
	case "||":
		return &values.Boolean{left.Value || right.Value, l.GetTok()}
	case "&&":
		return &values.Boolean{left.Value && right.Value, l.GetTok()}
	default:
		message.RaiseError("Operator", "This operator is not defined on booleans", left.GetTok())
	}
	return NIL
}

func evalStrInfix(l, r values.Value, op string) values.Value {
	if l.Type() != r.Type() {
		message.RaiseError("Type", "Invalid type on right side", r.GetTok())
	}
	left := l.(*values.String)
	right := r.(*values.String)

	switch op {
	case "+":
		return &values.String{left.Value + right.Value, l.GetTok()}
	case "==":
		return &values.Boolean{left.Value == right.Value, l.GetTok()}
	case "!=":
		return &values.Boolean{left.Value != right.Value, l.GetTok()}
	case "<=":
		return &values.Boolean{left.Value <= right.Value, l.GetTok()}
	case ">=":
		return &values.Boolean{left.Value >= right.Value, l.GetTok()}
	case "<":
		return &values.Boolean{left.Value < right.Value, l.GetTok()}
	case ">":
		return &values.Boolean{left.Value > right.Value, l.GetTok()}
	default:
		message.RaiseError("Operator", "Cannot eval string with this operator '"+op+"'", l.GetTok())
	}
	return NIL
}

func evalListInfix(l, r values.Value, op string) values.Value {
	if l.Type() != r.Type() {
		message.RaiseError("Type", "Invalid type on right side", r.GetTok())
	}
	left := l.(*values.List)
	right := r.(*values.List)

	switch op {
	case "+":
		return &values.List{append(left.Values, right.Values...), l.GetTok()}
	default:
		message.RaiseError("Operator", "Cannot eval list with this operator", l.GetTok())
	}
	return NIL
}

func evalObjInfix(n *ast.InfixExpr, consts *values.SymTab) values.Value {
	obj := EvalExpr(n.Left, consts)

	if obj.Type() != values.OBJ {
		message.RaiseError("Type", "Expected and object on the left side of '.'", n.GetTok())
	}

	return EvalExpr(n.Right, obj.(*values.Object).Consts)
}
