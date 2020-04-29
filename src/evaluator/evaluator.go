package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/tokens"
	"ludwig/src/values"
)

var (
	TOK = tokens.Token{"", 0, 0, "", tokens.EOL}
	NIL = &values.List{[]values.Value{}, TOK}
)

/* You may notice that in the parser we mapped the tokens to
 * their proper parsing functions. We did not use a similar
 * tactic here because the 'n := n.(type)' statement is only
 * possible in a 'switch' statement. This allows us to pass it
 * into the function as the proper type, and not the generic
 * 'ast.Node' type.
 */
func EvalExpr(n ast.Node, consts *values.SymTab, log *message.Log) values.Value {
	switch n := n.(type) {
	case *ast.Number:
		return evalNum(n)
	case *ast.String:
		return evalStr(n)
	case *ast.Boolean:
		return evalBool(n)
	case *ast.Nil:
		return evalNil(n)
	case *ast.Identifier:
		return evalIdent(n, consts, log)
	case *ast.List:
		return evalList(n, consts, log)
	case *ast.Index:
		return evalIndex(n, consts, log)
	case *ast.InfixExpr:
		return evalInfix(n, consts, log)
	case *ast.Block:
		return evalBlock(n, consts, log)
	case *ast.IfEl:
		return evalIfEl(n, consts, log)
	case *ast.Function:
		return evalFunc(n, consts, log)
	case *ast.Call:
		return evalCall(n, consts, log)
	case *ast.PrefixExpr:
		return evalPrefix(n, consts, log)
	case *ast.Struct:
		return evalStruct(n, consts, log)
	case *ast.Import:
		return evalImport(n, consts, log)
	case *ast.For:
		return evalForLoop(n, consts, log)
	case *ast.While:
		return evalWhileLoop(n, consts, log)
	case *ast.Slice:
		return evalSlice(n, consts, log)
	case *ast.TypeIdent:
		return evalTypeIdent(n)
	default:
		message.RuntimeErr("Eval", "Cannot evaluate this expression", n.GetTok(), log)
	}

	return NIL
}
