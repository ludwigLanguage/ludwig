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
	numberOfDefinitions int
}

func NewST() *SymTab {
	vals := make(map[string]Symbol)
	return &SymTab{vals, 0}
}

func (s *SymTab) Define(id string) Symbol {
	symbol := Symbol{GLOBAL_SCOPE, s.numberOfDefinitions}
	s.values[id] = symbol
	s.numberOfDefinitions++
	return symbol
}

func (s *SymTab) Resolve(id string) (Symbol, bool) {
	symbol, ok := s.values[id]
	return symbol, ok
}
