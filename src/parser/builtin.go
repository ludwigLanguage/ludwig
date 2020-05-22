package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

func (p *Parser) parseBuiltin() ast.Node {
	tok := p.lxr.CurTok
	name := p.lxr.CurTok.Value
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.LPAREN {
		p.raiseError("Syntax", "Expected '(' got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	args := p.getCallArgs()
	return ast.Builtin{tok, name, args}
}
