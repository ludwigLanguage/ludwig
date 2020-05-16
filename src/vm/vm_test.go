package vm

import (
	"fmt"
	"ludwig/src/ast"
	"ludwig/src/compiler"
	"ludwig/src/lexer"
	"ludwig/src/parser"
	"ludwig/src/source"
	"ludwig/src/values"
	"testing"
)

type vmTest struct {
	input    string
	expected interface{}
}

func TestNumMath(t *testing.T) {
	tests := []vmTest{
		{"program main\n1", 1.0},
		{"program main\n2", 2.0},
		{"program main\n1 + 2", 3.0},
	}

	runVmTest(t, tests)
}

func runVmTest(t *testing.T, tests []vmTest) {
	t.Helper()

	for iter, test := range tests {
		tree := parse(test.input)

		comp := compiler.New()
		comp.Compile(tree)

		compiled := comp.GetCompiled()
		fmt.Println(compiled)

		vm := New(compiled)
		vm.Run()

		stackElement := vm.LastPopped()
		err := testExpectedVal(t, test.expected, stackElement)
		if err != nil {
			t.Errorf("%s", err)
		} else {
			fmt.Printf("Passed test %v\n", iter+1)
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

func testExpectedVal(t *testing.T, expected interface{}, got values.Value) error {
	t.Helper()

	switch expected := expected.(type) {
	case float64:
		return testNumObj(float64(expected), got)
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
