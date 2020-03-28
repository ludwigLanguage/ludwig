package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

// Syntax: [<value, <value>]
func (p *Parser) parseList() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp() //move over '[' token

	entries := []ast.Node{}

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		entries = p.parseArgs()
	}

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax",
			"Expected ']' at the end of a list not '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return &ast.List{entries, tok}

}

/////////////////////////////////////////////////

//Syntax: <value>[<number>]
func (p *Parser) parseIndex(src ast.Node) ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	val := p.parseExpr(0)

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax",
			"Expected ']' after end of index but got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return &ast.Index{src, val, tok}
}
