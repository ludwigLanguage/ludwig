package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)

func (p *Parser) ParseProgram() {
	p.skipWhitespace()

	tok := p.getProgramTok()
	programExprs := p.getProgramExprs()
	exprs := p.castAsAssignments(programExprs)

	p.Tree = ast.Program{exprs, tok}
}

func (p *Parser) getProgramTok() tokens.Token {
	if p.lxr.CurTok.Alias != tokens.PROG {
		p.raiseError("Syntax", "Expected 'program' statement at the opening of the file")
	}
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	return tok
}

func (p *Parser) getProgramExprs() []ast.Node {
	exprs := []ast.Node{}
	for !p.lxr.IsDone() {
		expr := p.parseExpr(0)
		exprs = append(exprs, expr)
	}
	//TODO: assignments := p.castAsAssignments(exprs)
	return exprs
}

func (p *Parser) castAsAssignments(nodes []ast.Node) []ast.InfixExpr {
	assignments := []ast.InfixExpr{}
	for _, i := range nodes {
		if i.Type() != ast.INFIX {
			tok := i.GetTok()
			p.raiseErrorWithTok("Procedural", "Cannot have non-assignment statements outside function body", tok)
		}

		if i.(ast.InfixExpr).Op != "=" {
			tok := i.GetTok()
			p.raiseErrorWithTok("Procedural", "Cannot have non-assignment statements outside function body", tok)
		}

		assignments = append(assignments, i.(ast.InfixExpr))
	}

	return assignments
}
