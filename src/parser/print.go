package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

func (p *Parser) parsePrint() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	args := []ast.Node{}
	for p.lxr.CurTok.Alias != tokens.EOL && p.lxr.CurTok.Alias != tokens.EOF {
		args = append(args, p.parseExpr(0))
	}

	return ast.Print{tok, args}
}
