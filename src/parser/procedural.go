package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

// Syntax: if <expr> <expr> OR if <expr> <expr> else <expr>
func (p *Parser) parseIfEl() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	cond := p.parseExpr(0)
	expr := p.parseExpr(0)

	ifel := ast.IfEl{}
	ifel.Tok = tok
	ifel.Cond = cond
	ifel.Do = expr

	p.skipWhitespace()
	if p.lxr.CurTok.Alias == tokens.ELSE {
		p.lxr.MoveUp()
		ifel.ElseExpr = p.parseExpr(0)
	} else {
		ifel.ElseExpr = ast.Block{[]ast.Node{}, false, p.lxr.CurTok}
	}

	return ifel
}

//Syntax: for <ident>, <ident> in <expr> <expr>
func (p *Parser) parseForLoop() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.IDENT {
		p.raiseError("Syntax", "Expected an identifier, got '"+p.lxr.CurTok.Value+"'")
	}
	indexNumIdent := p.parseExpr(0).(*ast.Identifier)

	if p.lxr.CurTok.Alias != tokens.COMMA {
		p.raiseError("Syntax", "Expected ',' got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.IDENT {
		p.raiseError("Syntax", "Expected an identifier, got '"+p.lxr.CurTok.Value+"'")
	}
	indexIdent := p.parseExpr(0).(*ast.Identifier)

	if p.lxr.CurTok.Alias != tokens.IN {
		p.raiseError("Syntax", "Expected 'in' got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	list := p.parseExpr(0)
	doExpr := p.parseExpr(0)

	isScoped := false
	if doExpr.Type() == ast.BLOCK {
		isScoped = doExpr.(ast.Block).IsScoped
	}

	return ast.For{indexNumIdent, indexIdent, list, doExpr, isScoped, tok}
}

///////////////////////////////////////////////////////
// Syntax: while <condition> <expr>
func (p *Parser) parseWhileLoop() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	cond := p.parseExpr(0)
	do := p.parseExpr(0)
	isScoped := false

	if do.Type() == ast.BLOCK {
		isScoped = do.(ast.Block).IsScoped
	}

	return ast.While{cond, do, isScoped, tok}
}
