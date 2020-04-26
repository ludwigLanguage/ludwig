package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

//Syntax: func(<args>) <expr>
func (p *Parser) parseFunction() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.LPAREN {
		p.raiseError("Syntax", "Expected '(' before function arguments")
	}
	p.lxr.MoveUp()

	args := []ast.Node{}

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		args = p.parseArgs()
	}

	for _, i := range args {
		if i.Type() != ast.IDENT {
			p.raiseError("Syntax", "Arguments must be identifiers")
		}
	}

	argv := []*ast.Identifier{}
	for _, i := range args {
		argv = append(argv, i.(*ast.Identifier))
	}

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' after function arguments")
	}
	p.lxr.MoveUp()

	isVariadic := false
	if p.lxr.CurTok.Alias == tokens.DOT {
		p.lxr.MoveUp()
		if p.lxr.CurTok.Alias != tokens.DOT {
			p.raiseError("Syntax", "Expected '.'")
		}
		p.lxr.MoveUp()
		if p.lxr.CurTok.Alias != tokens.DOT {
			p.raiseError("Syntax", "Expected '.'")
		}
		p.lxr.MoveUp()
		isVariadic = true
	}

	if isVariadic && (len(argv) == 0) {
		p.raiseError("Syntax", "Cannot have variadic function with no arguments")
	}

	expr := p.parseExpr(0)

	return &ast.Function{argv, expr, isVariadic, tok}
}

/////////////////////////////////////////////////
//Syntax: <function>(<arguments>)
func (p *Parser) parseCall(callVal ast.Node) ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	args := []ast.Node{}
	if p.lxr.CurTok.Alias != tokens.RPAREN {
		args = p.parseArgs()
	}

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' at the end of a call")
	}
	p.lxr.MoveUp()

	return &ast.Call{callVal, args, tok}
}
