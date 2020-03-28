/* This interface will allow us to pass around all the ast nodes
 * in the evaluator as though they are all the same type
 */
package ast

import (
	"ludwig/src/tokens"
)

type Node interface {
	PrintAll(string)
	GetTok() tokens.Token
}
