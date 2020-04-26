package parser

import (
	"ludwig/src/ast"
)

//Syntax: struct <expr>
func (p *Parser) parseStruct() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	body := p.parseExpr(0)
	if body.Type() == "<block>" {
		body.(*ast.Block).IsScoped = false
	}

	return &ast.Struct{tok, body}
}
