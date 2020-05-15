package compiler

import (
	"fmt"
	"ludwig/src/ast"
	"ludwig/src/bytecode"
	"ludwig/src/lexer"
	"ludwig/src/parser"
	"ludwig/src/source"
	"ludwig/src/values"
	"testing"
)

type compilerTest struct {
	input                string
	expectedPool         []interface{}
	expectedInstructions []bytecode.Instructions
}

func TestIntMath(t *testing.T) {
	tests := []compilerTest{
		{
			input:        "program main\n1 + 2",
			expectedPool: []interface{}{1.0, 2.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.ADD),
			},
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTest) {
	t.Helper()

	for _, test := range tests {
		tree := parse(test.input)

		compiler := New()
		compiler.Compile(tree)

		program := compiler.GetCompiled()

		err := testInstructions(test.expectedInstructions, program.Instructions)
		if err != nil {
			t.Fatalf("Test Instructions Failed: %s", err)
		}

		err = testPool(t, test.expectedPool, program.Pool)
		if err != nil {
			t.Fatalf("Test Pool Failed: %s", err)
		}
	}
}

func parse(input string) *ast.Program {
	src := source.NewWithStr(input, "TEST CASE")
	lex := lexer.New(src)
	prs := parser.New(lex)
	prs.ParseProgram()

	return prs.Tree.(*ast.Program)
}

func testInstructions(expected []bytecode.Instructions, got bytecode.Instructions) error {
	wanted := concatInstructions(expected)

	if len(got) != len(wanted) {
		return fmt.Errorf("Incorrect Instruction Length.\nWant: %q\nGot: %q", wanted, got)
	}

	for i, instruction := range wanted {
		if got[i] != instruction {
			return fmt.Errorf("Incorrect Instruction At %d.\nWant: %q\nGot: %q", i, wanted, got)
		}
	}

	return nil
}

func concatInstructions(instructions []bytecode.Instructions) bytecode.Instructions {
	out := bytecode.Instructions{}

	for _, instruction := range instructions {
		out = append(out, instruction...)
	}

	return out
}

func testPool(t *testing.T, expected []interface{}, got []values.Value) error {
	if len(expected) != len(got) {
		return fmt.Errorf("Wrong Number of Constants in Pool.\nWant: %q\nGot: %q", len(expected), len(got))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case float64:
			return testNumObj(constant, got[i])
		}
	}

	return nil
}

func testNumObj(expected float64, got values.Value) error {
	result, ok := got.(*values.Number)
	if !ok {
		return fmt.Errorf("Expected Number Value.\nGot: %T (%+v)", got, got)
	}

	if result.Value != expected {
		return fmt.Errorf("Number Has Incorrect Value.\nGot: %v\nWant: %v", result.Value, expected)
	}

	return nil
}
