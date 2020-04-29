package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"math"
)

func evalInfix(n *ast.InfixExpr, consts *values.SymTab, log *message.Log) values.Value {

	if n.Op == "=" {
		return evalAssignment(n, consts, log)
	} else if n.Op == "." {
		return evalObjInfix(n, consts, log)
	}

	leftVal := EvalExpr(n.Left, consts, log)

	rightVal := EvalExpr(n.Right, consts, log)
	if leftVal.Type() == values.NIL || rightVal.Type() == values.NIL {
		isTheSame := leftVal.Type() == rightVal.Type()
		if n.Op == "==" {
			return &values.Boolean{isTheSame, n.GetTok()}
		} else if n.Op == "!=" {
			return &values.Boolean{!isTheSame, n.GetTok()}
		} else {
			message.RuntimeErr("Type", "Cannot evaluate an infix expression for these types", n.GetTok(), log)
		}
	}

	switch leftVal.Type() {
	case values.NUM:
		return evalNumInfix(leftVal, rightVal, n.Op, log)
	case values.BOOL:
		return evalBoolInfix(leftVal, rightVal, n.Op, log)
	case values.STR:
		return evalStrInfix(leftVal, rightVal, n.Op, log)
	case values.LIST:
		return evalListInfix(leftVal, rightVal, n.Op, log)
	default:
		message.RuntimeErr("Type", "Cannot evaluate an infix expression for these types", n.GetTok(), log)
	}
	return NIL
}

func evalNumInfix(l, r values.Value, op string, log *message.Log) values.Value {
	if l.Type() != r.Type() {
		message.RuntimeErr("Type", "Invalid right side type", r.GetTok(), log)
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
		message.RuntimeErr("Operator", "This operator is not defined on numbers", left.GetTok(), log)
	}

	return NIL
}

func evalBoolInfix(l, r values.Value, op string, log *message.Log) values.Value {
	if l.Type() != r.Type() {
		message.RuntimeErr("Type", "Invalid type on the right side", r.GetTok(), log)
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
		message.RuntimeErr("Operator", "This operator is not defined on booleans", left.GetTok(), log)
	}
	return NIL
}

func evalStrInfix(l, r values.Value, op string, log *message.Log) values.Value {
	if l.Type() != r.Type() {
		message.RuntimeErr("Type", "Invalid type on right side", r.GetTok(), log)
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
		message.RuntimeErr("Operator", "Cannot eval string with this operator '"+op+"'", l.GetTok(), log)
	}
	return NIL
}

func evalListInfix(l, r values.Value, op string, log *message.Log) values.Value {
	if l.Type() != r.Type() {
		message.RuntimeErr("Type", "Invalid type on right side", r.GetTok(), log)
	}
	left := l.(*values.List)
	right := r.(*values.List)

	switch op {
	case "+":
		return &values.List{append(left.Values, right.Values...), l.GetTok()}
	default:
		message.RuntimeErr("Operator", "Cannot eval list with this operator", l.GetTok(), log)
	}
	return NIL
}

func evalObjInfix(n *ast.InfixExpr, consts *values.SymTab, log *message.Log) values.Value {
	obj := EvalExpr(n.Left, consts, log)

	if obj.Type() != values.OBJ {
		message.RuntimeErr("Type", "Expected and object on the left side of '.'", n.GetTok(), log)
	}

	return EvalExpr(n.Right, obj.(*values.Object).Consts, log)
}
