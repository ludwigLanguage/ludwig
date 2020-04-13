package values

import (
	"fmt"
	"ludwig/src/ast"
)

type SymTab struct {
	values map[string]Value
}

func NewSymTab() *SymTab {
	v := make(map[string]Value)
	return &SymTab{v}
}

func (s *SymTab) SetVal(name string, v Value) {
	s.values[name] = v
}

func (s *SymTab) GetVal(name string) Value {
	return s.values[name]
}

func (s *SymTab) PrintAll() {
	fmt.Println("--------")
	for k, v := range s.values {
		fmt.Println(k, v.Stringify())
	}
}

func (s *SymTab) AddValsFrom(newSt *SymTab) {
	for k, v := range newSt.values {
		s.values[k] = v
	}
}

func (s *SymTab) AddValsFromExcept(newSt *SymTab, names []*ast.Identifier) {
	for k, v := range newSt.values {
		isInNames := false
		for _, i := range names {
			if i.Value == k {
				isInNames = true
			}
		}

		if !isInNames {
			s.values[k] = v
		}
	}
}