/* This file maps various symbols and keywords to the relevent "tokent aliases"
 * The lexer uses these maps to determine which token aliasess go with what
 * symbols
 */
package lexer

import (
	"ludwig/src/tokens"
)

var singleChar = map[string]byte{
	string(0): tokens.EOF,
	"\n":      tokens.EOL,
	";":       tokens.EOL,
	"!":       tokens.POP,
	"=":       tokens.OP5,
	"[":       tokens.LBRACK,
	"]":       tokens.RBRACK,
	"{":       tokens.LCURL,
	"}":       tokens.RCURL,
	"(":       tokens.LPAREN,
	")":       tokens.RPAREN,
	"+":       tokens.OP1,
	"-":       tokens.OP1,
	"*":       tokens.OP2,
	"/":       tokens.OP2,
	"^":       tokens.OP3,
	"<":       tokens.OP4,
	">":       tokens.OP4,
	":":       tokens.COLON,
	",":       tokens.COMMA,
	".":       tokens.DOT,
}

var doubleChar = map[string]byte{
	"==": tokens.OP4,
	"!=": tokens.OP4,
	">=": tokens.OP4,
	"<=": tokens.OP4,
	"||": tokens.OP5,
	"&&": tokens.OP5,
}

var keywords = map[string]byte{
	"true":   tokens.BOOL,
	"false":  tokens.BOOL,
	"nil":    tokens.NIL,
	"if":     tokens.IF,
	"else":   tokens.ELSE,
	"func":   tokens.FN,
	"struct": tokens.STRUCT,
	"import": tokens.IMPORT,
	"do"	: tokens.DO,
	"end"	: tokens.END,
}
