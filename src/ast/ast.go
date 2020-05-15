/* This interface will allow us to pass around all the ast nodes
 * in the evaluator as though they are all the same type
 */
package ast

import (
	"ludwig/src/tokens"
)

type Node interface {
	PrintAll(string)
	Stringify(string) string
	GetTok() tokens.Token
	Type() byte
}

//TODO: Convert to bytes
const (
	FN byte = iota
	CALL
	LIST
	INDEX
	QUOTE
	PREFIX
	INFIX
	NUM
	IDENT
	STR
	NIL
	BOOL
	BLOCK
	IFEL
	IMPRT
	STRCT
	FOR
	WHILE
	SLICE
	T_IDENT
	PROG
	PACK
)
