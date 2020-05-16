package parser

import "ludwig/src/ast"

func (p *Parser) parseTypeIdent() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	return ast.TypeIdent{tok.Value, tok}
}
