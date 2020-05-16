package values

import (
	"strconv"
)

type String struct {
	Value string
}

func (s String) Stringify() string {
	return s.Value
}

func (s String) Type() byte {
	return STR
}

/////////////////////////////////////////////////

type Number struct {
	Value float64
}

func (n Number) Stringify() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

func (n Number) Type() byte {
	return NUM
}

/////////////////////////////////////////////////

type Boolean struct {
	Value bool
}

func (b Boolean) Stringify() string {
	return strconv.FormatBool(b.Value)
}

func (b Boolean) Type() byte {
	return BOOL
}

/////////////////////////////////////////

type Nil struct{}

func (n Nil) Stringify() string {
	return "nil"
}

func (n Nil) Type() byte {
	return NIL
}
