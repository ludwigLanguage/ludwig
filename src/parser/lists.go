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
		entries = p.parseCommaSeparatedList()
	}

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax",
			"Expected ']' at the end of a list not '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return ast.List{entries, tok}

}

/////////////////////////////////////////////////

//Syntax: <value>[<number>]
func (p *Parser) parseIndex(src ast.Node) ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias == tokens.COLON {
		val := ast.Number{0, src.GetTok()}
		return p.parseSlice(val, src)
	}

	val := p.parseExpr(0)

	if p.lxr.CurTok.Alias == tokens.COLON {
		return p.parseSlice(val, src)
	}

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax",
			"Expected ']' after end of index but got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return ast.Index{src, val, tok}
}

func (p *Parser) parseSlice(startVal, src ast.Node) ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias == tokens.RBRACK {
		p.lxr.MoveUp()
		return ast.Slice{src, startVal, nil, tok}
	}

	endVal := p.parseExpr(0)

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax",
			"Expected ']' after end of slice but got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return ast.Slice{src, startVal, endVal, tok}
}
