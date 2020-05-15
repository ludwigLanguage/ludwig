package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

//Syntax: func(<args>) <expr>
func (p *Parser) parseFunction() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	argv := p.parseFnArgs()

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' after function arguments")
	}
	p.lxr.MoveUp()

	isVariadic := p.determineIfFuncIsVariadic()
	if isVariadic {
		p.checkVariadicArgLen(argv)
	}

	expr := p.parseExpr(0)

	return &ast.Function{argv, expr, isVariadic, tok}
}

func (p *Parser) parseFnArgs() []*ast.Identifier {
	if p.lxr.CurTok.Alias != tokens.LPAREN {
		p.raiseError("Syntax", "Expected '(' before function arguments")
	}
	p.lxr.MoveUp()

	args := []ast.Node{}

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		args = p.parseCommaSeparatedList()
	}

	argv := p.ensureNodesAreIdents(args)

	return argv
}

func (p *Parser) ensureNodesAreIdents(nodes []ast.Node) []*ast.Identifier {
	argv := []*ast.Identifier{}

	for _, i := range nodes {
		if i.Type() != ast.IDENT {
			p.raiseError("Syntax", "Expected identifiers in function arguments")
		}

		argv = append(argv, i.(*ast.Identifier))
	}

	return argv
}

func (p *Parser) determineIfFuncIsVariadic() bool {
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

	return isVariadic
}

func (p *Parser) checkVariadicArgLen(args []*ast.Identifier) {
	if len(args) == 0 {
		p.raiseError("Syntax", "Cannot have variadic function with no arguments")
	}
}
