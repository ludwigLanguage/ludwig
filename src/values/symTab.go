package values

import (
	"fmt"
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
