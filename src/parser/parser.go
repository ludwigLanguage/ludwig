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
		tokens.NUM:         p.parseNum,
		tokens.STR:         p.parseStr,
		tokens.BOOL:        p.parseBool,
		tokens.IDENT:       p.parseIdent,
		tokens.LBRACK:      p.parseList,
		tokens.POP:         p.parsePrefix,
		tokens.OP1:         p.parsePrefix,
		tokens.DO:          p.parseBlock,
		tokens.NIL:         p.parseNil,
		tokens.LPAREN:      p.parseLParen,
		tokens.IF:          p.parseIfEl,
		tokens.FN:          p.parseFunction,
		tokens.FOR:         p.parseForLoop,
		tokens.WHILE:       p.parseWhileLoop,
		tokens.TYPE_INDENT: p.parseTypeIdent,
		tokens.IMPORT:      p.parseImport,
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

func (p *Parser) raiseErrorWithTok(errorName, errorMessage string, token tokens.Token) {
	message.RaiseError(errorName, errorMessage, token)
}

func (p *Parser) parseExpr(prec int) ast.Node {
	p.skipWhitespace()

	preFn := p.prefixParseFns[p.lxr.CurTok.Alias]
	if preFn == nil {
		p.raiseError("Syntax",
			"Cannot parse expression starting with: '"+p.lxr.CurTok.Value+"'")
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

func (p *Parser) parseLParen() ast.Node {
	p.lxr.MoveUp()
	expr := p.parseExpr(0)

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return expr
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
