package parser

import (
	"ludwig/src/ast"
)
//Syntax: <operator><expression>
func (p *Parser) parsePrefix() ast.Node {
	tok := p.lxr.CurTok
	op := p.lxr.CurTok.Value
	p.lxr.MoveUp()

	/* The precedence on prefix operators must fall in between
	 * LPAREN (object and function calls), and other types of binary
	 * operators (such as '+' or '/') so that prefix operators are evaluated
	 * before the other math, but after function calls 
	 */
	expr := p.parseExpr(6)

	return &ast.PrefixExpr{expr, op, tok}
}

/////////////////////////////////////////////////
//Syntax: <expr> <operator> <expr>
func (p *Parser) parseInfix(left ast.Node) ast.Node {
	op := p.lxr.CurTok.Value
	tok := p.lxr.CurTok

	prec := precedence[p.lxr.CurTok.Alias]
	p.lxr.MoveUp()
	right := p.parseExpr(prec)

	return &ast.InfixExpr{left, right, op, tok}
}

/////////////////////////////////////////////////
