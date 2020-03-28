package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/values"
	"ludwig/src/tokens"
	"ludwig/src/source"
	"ludwig/src/lexer"
	"ludwig/src/parser"
)

func evalBlock(n *ast.Block, consts *values.SymTab) values.Value {
	var newC *values.SymTab
	isStruct := isStructSymTab(n.GetTok(), consts)
	if n.IsScoped && !isStruct {
		newC = newSymTabCopy(consts)
	} else {
		newC = consts
	}

	var val values.Value
	for _, i := range n.Body {
		val = EvalExpr(i, newC)
	}

	if val == nil {
		return NIL
	}

	return val
}

/* If a scope contains a self object, then it must be
 * a struct
 */
func isStructSymTab(tok tokens.Token, consts *values.SymTab) bool {
	selfObj := consts.GetVal("self")
	return selfObj != nil
}

func newSymTabCopy(consts *values.SymTab) *values.SymTab {

	newC := values.NewSymTab()
	newC.AddValsFrom(consts)

	return newC
}

func evalIfEl(n *ast.IfEl, consts *values.SymTab) values.Value {
	cond := EvalExpr(n.Cond, consts)

	if cond.Type() != values.BOOL {
		message.RaiseError("Type", "This expression must evaluate to a boolean", n.GetTok())
	}

	if cond.(*values.Boolean).Value {
		return EvalExpr(n.Do, consts)
	}

	return EvalExpr(n.ElseExpr, consts)
}

func evalImport(n *ast.Import, consts *values.SymTab) values.Value {
	fileVal := EvalExpr(n.Filename, consts)

	if fileVal.Type() != values.STR {
		message.RaiseError("Type", "Import function can only recieve strings", n.GetTok())
	}
	filename := fileVal.(*values.String).Value


	src := source.New(filename)
	lxr := lexer.New(src)
	prs := parser.New(lxr)
	prs.ParseProgram()
	
	symTabForObj := values.NewSymTab()
	EvalExpr(prs.Tree, symTabForObj)

	return &values.Object {symTabForObj, n.GetTok()}
}