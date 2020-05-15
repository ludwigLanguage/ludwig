package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

/* Syntax:
 * package <identifier>
 * public:
 *	<public_contents>
 * private:
 *   <private_contents>
 */
func (p *Parser) ParsePackage() ast.Node {

	tok := p.getPackageTok()
	id := p.getProgramOrPackageId()
	pubExprs := p.getPublicExprs()
	privExprs := []*ast.InfixExpr{}

	if !p.lxr.IsDone() {
		privExprs = p.getPrivateExprs()
	}

	return &ast.Package{id, pubExprs, privExprs, tok}
}

func (p *Parser) getPackageTok() tokens.Token {
	p.skipWhitespace()

	if p.lxr.CurTok.Alias != tokens.PACK {
		p.raiseError("Syntax", "Expected 'package' at the beginning of the file")
	}
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	return tok
}

func (p *Parser) getPublicExprs() []*ast.InfixExpr {
	p.skipWhitespace()

	if p.lxr.CurTok.Alias != tokens.PUB {
		p.raiseError("Syntax", "Expected 'public' after package opening")
	}
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.COLON {
		p.raiseError("Syntax", "Expected ':' after 'public'")
	}
	p.lxr.MoveUp()

	exprs := []ast.Node{}
	for p.lxr.CurTok.Alias != tokens.PRIV && !p.lxr.IsDone() {
		exprs = append(exprs, p.parseExpr(0))
		p.skipWhitespace()
	}

	return p.castAsAssignments(exprs)
}

func (p *Parser) getPrivateExprs() []*ast.InfixExpr {
	p.skipWhitespace()

	if p.lxr.CurTok.Alias != tokens.PRIV {
		p.raiseError("Syntax", "Expected 'private' after public declarations")
	}
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.COLON {
		p.raiseError("Syntax", "Expected ':' after 'private'")
	}
	p.lxr.MoveUp()

	exprs := []ast.Node{}
	for !p.lxr.IsDone() {
		exprs = append(exprs, p.parseExpr(0))
	}

	return p.castAsAssignments(exprs)
}
