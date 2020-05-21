package compiler

type symbol struct {
	Index int
	Scope int
}

type SymTab struct {
	values              map[string]symbol
	NumberOfDefinitions int
	previousStates      []map[string]symbol
	numOfSavedStates    int
	curScope            int
}

func NewST() *SymTab {
	vals := make(map[string]symbol)
	previous := make([]map[string]symbol, 32)
	return &SymTab{vals, 0, previous, 0, 0}
}

func (s *SymTab) Define(id string) int {

	/* If we are in the same scope the varible was
	 * defined in, then we are allowed to reuse it's
	 * place in memory safely
	 */
	sym, ok := s.values[id]
	if ok {
		if sym.Scope == s.curScope {
			return sym.Index
		}
	}

	s.values[id] = symbol{s.NumberOfDefinitions, s.curScope}
	s.NumberOfDefinitions++
	return s.NumberOfDefinitions - 1
}

func (s *SymTab) Resolve(id string) (int, bool) {
	symbol, ok := s.values[id]
	return symbol.Index, ok
}

func (s *SymTab) SaveState() {
	saveMap := map[string]symbol{}

	for k, v := range s.values {
		saveMap[k] = v
	}

	if s.numOfSavedStates >= cap(s.previousStates) {
		newPrevious := make([]map[string]symbol, len(s.previousStates)*2)
		for j, i := range s.previousStates {
			newPrevious[j] = i
		}
		s.previousStates = newPrevious
	}

	s.previousStates[s.numOfSavedStates] = saveMap
	s.numOfSavedStates++
	s.curScope++
}

func (s *SymTab) Revert() {
	s.numOfSavedStates--
	s.values = s.previousStates[s.numOfSavedStates]
	s.curScope--
}
