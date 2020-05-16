package values

type Value interface {
	Stringify() string
	Type() byte
}
