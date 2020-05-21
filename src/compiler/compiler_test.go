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
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n1-2",
			expectedPool: []interface{}{1.0, 2.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.SUB),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n1*2",
			expectedPool: []interface{}{1.0, 2.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.MULT),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n1/2",
			expectedPool: []interface{}{1.0, 2.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.DIV),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n1^2",
			expectedPool: []interface{}{1.0, 2.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.POW),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\ntrue",
			expectedPool: []interface{}{true},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n!true",
			expectedPool: []interface{}{true},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.NOT),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},

		{
			input:        "program main\n-10",
			expectedPool: []interface{}{10.0},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.NEGATIVE),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestIfEl(t *testing.T) {
	tests := []compilerTest{
		{
			input: `program main
			if true
				10
			333`,
			expectedPool: []interface{}{true, 10, nil, 333},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				//Jumps to byte we want, adjusts for different instruction widths
				bytecode.MakeInstruction(bytecode.JUMPNT, 12),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.JUMP, 15),
				bytecode.MakeInstruction(bytecode.LOADCONST, 2), //Empty Else Branch
				bytecode.MakeInstruction(bytecode.POP),
				bytecode.MakeInstruction(bytecode.LOADCONST, 3),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestGlobals(t *testing.T) {
	tests := []compilerTest{
		{
			input: `program main
			a = 10
			a`,
			expectedPool: []interface{}{10},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.SAVEV),
				bytecode.MakeInstruction(bytecode.POP), //Pop off value made by a = 10 expression
				bytecode.MakeInstruction(bytecode.GETV, 0),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestList(t *testing.T) {
	tests := []compilerTest{
		{
			input:        "program main; [1, 2, 3]",
			expectedPool: []interface{}{1, 2, 3},
			expectedInstructions: []bytecode.Instructions{
				bytecode.MakeInstruction(bytecode.LOADCONST, 0),
				bytecode.MakeInstruction(bytecode.LOADCONST, 1),
				bytecode.MakeInstruction(bytecode.LOADCONST, 2),
				bytecode.MakeInstruction(bytecode.BUILDLIST, 3),
				bytecode.MakeInstruction(bytecode.POP),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestSymTab(t *testing.T) {
	expected := map[string]int{
		"x": 0,
		"y": 1,
	}

	symtab := NewST()
	x := symtab.Define("x")
	if x != expected["x"] {
		t.Errorf("Did not get expected value binded\nGot=%v", x)
	}

	y := symtab.Define("y")
	if y != expected["y"] {
		t.Errorf("Did not get expected value binded")
	}
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

func parse(input string) ast.Program {
	src := source.NewWithStr(input, "TEST CASE")
	lex := lexer.New(src)
	prs := parser.New(lex)
	prs.ParseProgram()

	return prs.Tree.(ast.Program)
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
		case bool:
			return testBoolObj(constant, got[i])
		}
	}

	return nil
}

func testNumObj(expected float64, got values.Value) error {
	result, ok := got.(values.Number)
	if !ok {
		return fmt.Errorf("Expected Number Value.\nGot: %T (%+v)", got, got)
	}

	if result.Value != expected {
		return fmt.Errorf("Number Has Incorrect Value.\nGot: %v\nWant: %v", result.Value, expected)
	}

	return nil
}

func testBoolObj(expected bool, got values.Value) error {
	result, ok := got.(values.Boolean)
	if !ok {
		return fmt.Errorf("Expected boolean value\nGot: %T (%+v)", got, got)
	}

	if result.Value != expected {
		return fmt.Errorf("Boolean has incorrect value.\nGot: %v\nWant: %v", result.Value, expected)
	}

	return nil
}
