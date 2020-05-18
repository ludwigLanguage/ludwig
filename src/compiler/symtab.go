package compiler

const (
	GLOBAL_SCOPE byte = iota
)

type Symbol struct {
	Scope byte
	Index int
}

type SymTab struct {
	values              map[string]Symbol
	NumberOfDefinitions int
}

func NewST() *SymTab {
	vals := make(map[string]Symbol)
	return &SymTab{vals, 0}
}

func (s *SymTab) Define(id string) Symbol {
	symbol := Symbol{GLOBAL_SCOPE, s.NumberOfDefinitions}
	s.values[id] = symbol
	s.NumberOfDefinitions++
	return symbol
}

func (s *SymTab) Resolve(id string) (Symbol, bool) {
	symbol, ok := s.values[id]
	return symbol, ok
}

func (s *SymTab) ClearDefsBackTo(num int) {
	for k, v := range s.values {
		if v.Index >= num {
			delete(s.values, k)
		}
	}
}
