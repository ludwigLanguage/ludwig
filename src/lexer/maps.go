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
	"do":     tokens.DO,
	"end":    tokens.END,
	"for":    tokens.FOR,
	"in":     tokens.IN,
	"while":  tokens.WHILE,

	"_num":     tokens.TYPE_INDENT,
	"_str":     tokens.TYPE_INDENT,
	"_bool":    tokens.TYPE_INDENT,
	"_nil":     tokens.TYPE_INDENT,
	"_list":    tokens.TYPE_INDENT,
	"_func":    tokens.TYPE_INDENT,
	"_struct":  tokens.TYPE_INDENT,
	"_object":  tokens.TYPE_INDENT,
	"_builtin": tokens.TYPE_INDENT,
	"_type":    tokens.TYPE_INDENT,
}
