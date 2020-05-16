package values

/////////////////////////////////////////////////

type Builtin struct {
	Fn func([]Value) Value
}

func (b Builtin) Stringify() string {
	return "builtinFn()"
}

func (b Builtin) Type() byte {
	return BUILTIN
}
