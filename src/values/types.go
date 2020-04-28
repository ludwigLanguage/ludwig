package values

import "ludwig/src/tokens"

const (
	NUM        string = "_num"
	STR        string = "_str"
	BOOL       string = "_bool"
	NIL        string = "_nil"
	LIST       string = "_list"
	FUNC       string = "_func"
	STRUCT     string = "_struct"
	OBJ        string = "_object"
	BUILTIN    string = "_builtin"
	TYPE_IDENT string = "_type"
)

type TypeIdent struct {
	Value string
	Tok   tokens.Token
}

func (t *TypeIdent) Stringify() string {
	return t.Value
}

func (t *TypeIdent) Type() string {
	return TYPE_IDENT
}

func (t *TypeIdent) GetTok() tokens.Token {
	return t.Tok
}
