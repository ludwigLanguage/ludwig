package parser

import (
	"ludwig/src/ast"
	"strconv"
)

/* <number> OR <number>.<number>
 * Note: all numbers, whether they appear in the
 * source text as floats of intergers will be stored
 * as a float64.
 */
func (p *Parser) parseNum() ast.Node {
	v, err := strconv.ParseFloat(p.lxr.CurTok.Value, 64)

	if err != nil {
		p.raiseError("Syntax",
			"Could not parse number '"+p.lxr.CurTok.Value+"'")
	}

	n := ast.Number{v, p.lxr.CurTok}
	p.lxr.MoveUp()
	return n
}

/////////////////////////////////////////////////
// "<text>" or '<text>
func (p *Parser) parseStr() ast.Node {
	v := p.lxr.CurTok.Value
	n := ast.String{v, p.lxr.CurTok}

	p.lxr.MoveUp()
	return n
}

/////////////////////////////////////////////////
//Syntax: true OR false
func (p *Parser) parseBool() ast.Node {
	v, err := strconv.ParseBool(p.lxr.CurTok.Value)

	if err != nil {
		p.raiseError("Syntax", "Could not parse boolean '"+p.lxr.CurTok.Value+"'")
	}

	n := ast.Boolean{v, p.lxr.CurTok}
	p.lxr.MoveUp()
	return n
}

//////////////////////////////////////////////////
//Syntax: nil
func (p *Parser) parseNil() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	return ast.Nil{tok}
}

/////////////////////////////////////////////////
//Syntax: <ident>
func (p *Parser) parseIdent() ast.Node {
	v := p.lxr.CurTok.Value
	n := ast.Identifier{v, p.lxr.CurTok}
	p.lxr.MoveUp()

	return n
}
