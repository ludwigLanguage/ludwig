package values

import (
	"ludwig/src/tokens"
	"ludwig/src/ast"
)

type Value interface {
	Stringify() string
	Type() string
	GetTok() tokens.Token
	ConvertToAst() ast.Node
}
