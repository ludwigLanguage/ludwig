/*  This file contains the parser object which is the
 * object that actually does the dirty work of parsing
 * tokens into an ast
 */
package parser

import (
	"ludwig/src/ast"
	"ludwig/src/lexer"
	"ludwig/src/message"
	"ludwig/src/tokens"
)

/* This map associates various operators with the 
 * proper precedence. The lower the number, the lower the
 * operator will be placed in the tree
 */
var precedence = map[byte]int{
	tokens.OP1:    3,
	tokens.OP2:    4,
	tokens.OP3:    5,
	tokens.OP4:    2,
	tokens.LPAREN: 7,
	tokens.LBRACK: 7, 
	tokens.DOT:    8,
	tokens.OP5:    1,
}

type (
	prefixFn func() ast.Node
	infixFn  func(ast.Node) ast.Node
)

type Parser struct {
	lxr  *lexer.Lexer
	Tree ast.Node

	prefixParseFns map[byte]prefixFn
	infixParseFns  map[byte]infixFn
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{}
	p.lxr = lexer

	p.prefixParseFns = map[byte]prefixFn{
		tokens.NUM:    p.parseNum,
		tokens.STR:    p.parseStr,
		tokens.BOOL:   p.parseBool,
		tokens.IDENT:  p.parseIdent,
		tokens.LBRACK: p.parseList,
		tokens.POP:    p.parsePrefix,
		tokens.OP1:    p.parsePrefix,
		tokens.DO: 	   p.parseBlock,
		tokens.NIL:    p.parseNil,
		tokens.LCURL:  p.parseBlock,
		tokens.IF:     p.parseIfEl,
		tokens.FN:     p.parseFunction,
		tokens.STRUCT: p.parseStruct,
		tokens.IMPORT: p.parseImport,
		tokens.QUOTE: p.parseQuoteOrUnquote,
		tokens.UNQUOTE: p.parseQuoteOrUnquote,
	}

	p.infixParseFns = map[byte]infixFn{
		tokens.OP1:    p.parseInfix,
		tokens.OP2:    p.parseInfix,
		tokens.OP3:    p.parseInfix,
		tokens.OP4:    p.parseInfix,
		tokens.OP5:    p.parseInfix,
		tokens.DOT:    p.parseInfix,
		tokens.LBRACK: p.parseIndex,
		tokens.LPAREN: p.parseCall,
	}

	return p
}

func (p *Parser) raiseError(n, m string) {
	message.Error(p.lxr.CurTok.Filename, n, m,
		p.lxr.Src().LineNo, p.lxr.Src().ColumnNo)
}

func (p *Parser) ParseProgram() {
	p.Tree = p.parseExpr(0)
}

func (p *Parser) parseExpr(prec int) ast.Node {
	p.skipWhitespace()

	preFn := p.prefixParseFns[p.lxr.CurTok.Alias]
	if preFn == nil {
		p.raiseError("Syntax",
			"No prefix parse fn for '"+p.lxr.CurTok.Value+"'")
	}
	leftExpr := preFn()

	for p.notDoneParsingExpr(prec) {

		infFn := p.infixParseFns[p.lxr.CurTok.Alias]
		if infFn == nil {
			return leftExpr
		}

		leftExpr = infFn(leftExpr)
	}

	return leftExpr
}

func (p *Parser) notDoneParsingExpr(prec int) bool {
	return p.lxr.CurTok.Alias != tokens.EOL &&
		p.lxr.CurTok.Alias != tokens.EOF &&
		prec < precedence[p.lxr.CurTok.Alias]
}

func (p *Parser) skipWhitespace() {
	for p.lxr.CurTok.Alias == tokens.EOL {
		p.lxr.MoveUp()
	}
}
