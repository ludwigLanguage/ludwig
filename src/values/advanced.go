package values

import (
	"fmt"
	"ludwig/src/bytecode"
)

type Function struct {
	Instructions bytecode.Instructions
	NumOfArgs    int
}

func (f Function) Stringify() string {
	return fmt.Sprintf("function { %p }", f)
}

func (f Function) Type() byte {
	return FUNC
}
