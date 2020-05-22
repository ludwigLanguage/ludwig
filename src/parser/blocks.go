package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

/* Syntax: do { <exprs> } | do ( <expr> )
 * A scoped block is encased in brackets,
 * An un-scoped block is encased in parentheses
 */
func (p *Parser) parseBlock() ast.Node {
	tok := p.lxr.CurTok

	ending := p.getBlockEnding()
	isScoped := p.getScope()
	p.lxr.MoveUp()

	body := []ast.Node{}
	p.skipWhitespace()
	for p.notDoneParsingBlock(ending) {
		expr := p.parseExpr(0)
		body = append(body, expr)
	}
	p.lxr.MoveUp() //Move over closing bracket

	return ast.Block{body, isScoped, tok}
}

func (p *Parser) getBlockEnding() byte {
	if p.lxr.CurTok.Alias == tokens.LCURL {
		return tokens.RCURL
	} else if p.lxr.CurTok.Alias == tokens.LPAREN {
		return tokens.RPAREN
	} else {
		p.raiseError("Syntax", "Expected '{' or '(' after 'do' got: "+p.lxr.CurTok.Value)
	}
	return 0
}

func (p *Parser) getScope() bool {
	return p.lxr.CurTok.Alias == tokens.LCURL
}

func (p *Parser) notDoneParsingBlock(ending byte) bool {
	p.skipWhitespace()
	if p.lxr.CurTok.Alias == tokens.EOF {
		p.raiseError("Syntax", "Expected '}' or ')' before EOF")
	}

	return p.lxr.CurTok.Alias != ending
}
