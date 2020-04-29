package evaluator

import (
	"ludwig/src/ast"
	"ludwig/src/lexer"
	"ludwig/src/message"
	"ludwig/src/parser"
	"ludwig/src/source"
	"ludwig/src/values"
)

func evalBlock(n *ast.Block, consts *values.SymTab, log *message.Log) values.Value {
	var newC *values.SymTab

	if n.IsScoped {
		newC = newSymTabCopy(consts)
	} else {
		newC = consts
	}

	var val values.Value
	for _, i := range n.Body {
		val = EvalExpr(i, newC, log)
	}

	if val == nil {
		return NIL
	}

	return val
}

func newSymTabCopy(consts *values.SymTab) *values.SymTab {

	newC := values.NewSymTab()
	newC.AddValsFrom(consts)

	return newC
}

func evalIfEl(n *ast.IfEl, consts *values.SymTab, log *message.Log) values.Value {
	cond := EvalExpr(n.Cond, consts, log)

	if cond.Type() != values.BOOL {
		message.RuntimeErr("Type", "This expression must evaluate to a boolean", n.GetTok(), log)
	}

	if cond.(*values.Boolean).Value {
		return EvalExpr(n.Do, consts, log)
	}

	return EvalExpr(n.ElseExpr, consts, log)
}

func evalImport(n *ast.Import, consts *values.SymTab, log *message.Log) values.Value {
	fileVal := EvalExpr(n.Filename, consts, log)

	if fileVal.Type() != values.STR {
		message.RuntimeErr("Type", "Import function can only recieve strings", n.GetTok(), log)
	}
	filename := fileVal.(*values.String).Value

	src := source.New(filename)
	lxr := lexer.New(src)
	prs := parser.New(lxr)
	prs.ParseProgram()

	symTabForObj := values.NewSymTab()
	EvalExpr(prs.Tree, symTabForObj, log)

	return &values.Object{symTabForObj, n.GetTok()}
}

func evalForLoop(n *ast.For, consts *values.SymTab, log *message.Log) values.Value {
	rtrnList := []values.Value{}

	forSt := values.NewSymTab()
	forSt.AddValsFrom(consts)

	loopVal := EvalExpr(n.List, consts, log)

	if loopVal.Type() != values.LIST {
		message.RuntimeErr("Type", "Expected list got '"+loopVal.Type()+"'", loopVal.GetTok(), log)
	}

	//We can assume that the doExpr is a block if IsScoped is
	//true due to the way I wrote the parser method for for loops
	if n.IsScoped {
		n.DoExpr.(*ast.Block).IsScoped = false
	}

	loopLst := loopVal.(*values.List).Values
	for num, val := range loopLst {
		numVal := &values.Number{float64(num), n.GetTok()}

		//Add values to st
		id := n.IndexNumIdent.Value
		forSt.SetVal(id, numVal)

		id = n.IndexIdent.Value
		forSt.SetVal(id, val)
		//

		evaledExpr := EvalExpr(n.DoExpr, forSt, log)
		rtrnList = append(rtrnList, evaledExpr)

	}

	if !n.IsScoped {
		consts.AddValsFrom(forSt)
	}

	return &values.List{rtrnList, n.GetTok()}
}

func evalWhileLoop(n *ast.While, consts *values.SymTab, log *message.Log) values.Value {
	cond := EvalExpr(n.Cond, consts, log)

	if cond.Type() != values.BOOL {
		message.RuntimeErr("Type", "Conditional for while must yeild a boolean", n.GetTok(), log)
	}

	if n.IsScoped {
		n.Body.(*ast.Block).IsScoped = false
		consts = newSymTabCopy(consts)
	}

	rtrnList := []values.Value{}
	for cond.(*values.Boolean).Value {
		rtrnList = append(rtrnList, EvalExpr(n.Body, consts, log))

		cond = EvalExpr(n.Cond, consts, log)
	}

	return &values.List{rtrnList, n.GetTok()}
}
