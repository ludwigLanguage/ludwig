package values

const (
	NUM byte = iota
	STR
	BOOL
	NIL
	LIST
	FUNC
	STRUCT
	OBJ
	BUILTIN
	TYPE_IDENT
)

type TypeIdent struct {
	Value byte
}

func (t TypeIdent) Stringify() string {
	switch t.Value {
	case STR:
		return "<string>"
	case BOOL:
		return "<bool>"
	case NIL:
		return "<nil>"
	case LIST:
		return "<list>"
	case FUNC:
		return "<function>"
	case STRUCT:
		return "<struct>"
	case OBJ:
		return "<object>"
	case BUILTIN:
		return "<builtin>"
	case TYPE_IDENT:
		return "<type_identifier>"
	}

	return "<unknown_type>"
}

func (t TypeIdent) Type() byte {
	return TYPE_IDENT
}
