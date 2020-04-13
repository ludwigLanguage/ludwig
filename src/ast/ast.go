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
	Type() string
}

const (
	FN = "<function>"
	CALL = "<call>"
	LIST = "<list>"
	INDEX = "<index>"
	QUOTE = "<quote>"
	UNQUOTE = "<unquote>"
	PREFIX = "<prefixExpression>"
	INFIX = "<infixExpression>"
	NUM = "<number>"
	IDENT = "<identifier>"
	STR = "<string"
	NIL = "<nil>"
	BOOL = "<boolean>"
	BLOCK = "<block>"
	IFEL = "<ifElse>"
	IMPRT = "<import>"
	STRCT = "<structure>"
)