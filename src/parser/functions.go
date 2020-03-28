package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"

	"fmt"
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
		if fmt.Sprintf("%T", i) != "*ast.Identifier" {
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

	expr := p.parseExpr(0)

	return &ast.Function{argv, expr, tok}
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
