package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

func (p *Parser) parseQuoteOrUnquote() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.LPAREN {
		p.raiseError("Syntax", "Expected '(' not '" + p.lxr.CurTok.Value + "'")
	}
	p.lxr.MoveUp()

	expr := p.parseExpr(0)

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected '(' not '" + p.lxr.CurTok.Value + "'")
	}
	p.lxr.MoveUp()

	var rtrnVal ast.Node
	if tok.Alias == tokens.QUOTE {
		rtrnVal = &ast.Quote {expr, tok}
	} else {
		rtrnVal = &ast.UnQuote {expr, tok}
	}

	return rtrnVal
}

