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
func EvalExpr(n ast.Node, consts *values.SymTab) values.Value {
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
		return evalIdent(n, consts)
	case *ast.List:
		return evalList(n, consts)
	case *ast.Index:
		return evalIndex(n, consts)
	case *ast.InfixExpr:
		return evalInfix(n, consts)
	case *ast.Block:
		return evalBlock(n, consts)
	case *ast.IfEl:
		return evalIfEl(n, consts)
	case *ast.Function:
		return evalFunc(n, consts)
	case *ast.Call:
		return evalCall(n, consts)
	case *ast.PrefixExpr:
		return evalPrefix(n, consts)
	case *ast.Struct:
		return evalStruct(n, consts)
	case *ast.Import:
		return evalImport(n, consts)
	case *ast.For:
		return evalForLoop(n, consts)
	case *ast.While:
		return evalWhileLoop(n, consts)
	default:
		message.RaiseError("Eval", "Cannot evaluate this expression", n.GetTok())
	}

	return NIL
}
