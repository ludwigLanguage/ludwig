package compiler

const (
	GLOBAL_SCOPE byte = iota
)

type SymTab struct {
	values              map[string]int
	NumberOfDefinitions int
	previousStates      []map[string]int
	numOfSavedStates    int
}

func NewST() *SymTab {
	vals := make(map[string]int)
	previous := make([]map[string]int, 128)
	return &SymTab{vals, 0, previous, 0}
}

func (s *SymTab) Define(id string) int {
	s.values[id] = s.NumberOfDefinitions
	s.NumberOfDefinitions++
	return s.NumberOfDefinitions - 1
}

func (s *SymTab) Resolve(id string) (int, bool) {
	symbol, ok := s.values[id]

	return symbol, ok
}

func (s *SymTab) SaveState() bool {
	saveMap := map[string]int{}

	for k, v := range s.values {
		saveMap[k] = v
	}

	ok := true
	if s.numOfSavedStates >= 127 {
		ok = false
	}

	s.previousStates[s.numOfSavedStates] = saveMap
	s.numOfSavedStates++

	return ok
}

func (s *SymTab) Revert() {
	s.numOfSavedStates--
	s.values = s.previousStates[s.numOfSavedStates]
}
