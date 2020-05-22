package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

//Syntax: <function>(<arguments>)
func (p *Parser) parseCall(callVal ast.Node) ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	args := p.getCallArgs()

	return ast.Call{callVal, args, tok}
}

func (p *Parser) getCallArgs() []ast.Node {
	args := []ast.Node{}
	if p.lxr.CurTok.Alias != tokens.RPAREN {
		args = p.parseCommaSeparatedList()
	}
	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' at the end of a call")
	}
	p.lxr.MoveUp()
	return args
}
