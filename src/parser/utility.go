package parser

import (
	"ludwig/src/ast"
	"ludwig/src/tokens"
)
/* There is no ast node that is produced by this function,
 * and this function does not show in the ParseFns maps in
 * the Parser{} struct. This is instead used as a utility function
 * by various other parsing functions to parse items with the following
 * syntax: <expr>, <expr>, <expr>...
 * This function is used in the parsing of lists, call arguments,
 * and function declaration arguments.
 */
func (p *Parser) parseArgs() []ast.Node {
	lst := []ast.Node{p.parseExpr(0)}

	for p.lxr.CurTok.Alias == tokens.COMMA {
		p.lxr.MoveUp() //Move over comma
		lst = append(lst, p.parseExpr(0))
	}
	return lst
}
