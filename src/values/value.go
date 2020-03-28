package values

import (
	"ludwig/src/tokens"
)

type Value interface {
	Stringify() string
	Type() string
	GetTok() tokens.Token
}
