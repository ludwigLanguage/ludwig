package compiler

import (
	"ludwig/src/bytecode"
	"ludwig/src/values"
)

type CompiledProg struct {
	Instructions bytecode.Instructions
	Pool         []values.Value
}
