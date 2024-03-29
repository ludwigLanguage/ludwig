package values

import (
	"ludwig/src/tokens"
)

type List struct {
	Values []Value
	Tok    tokens.Token
}

func (l *List) Stringify() string {

	str := "["

	for _, i := range l.Values {
		str += i.Stringify() + ", "
	}

	if len(str) > 1 {
		str = str[:len(str)-2] //Get ride of trailing comma
	}

	return str + "]"
}

func (l *List) Type() string {
	return LIST
}

func (l *List) GetTok() tokens.Token {
	return l.Tok
}

/////////////////////////////////////////////////
