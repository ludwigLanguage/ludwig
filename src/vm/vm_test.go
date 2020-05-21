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
		{"program main\n1-2", -1.0},
		{"program main\n2*2", 4.0},
		{"program main\n9/3", 3.0},
		{"program main\n2^4", 16.0},
		{"program main\n-4", -4.0},
	}

	runVmTest(t, tests)
}

func TestStringConcat(t *testing.T) {
	tests := []vmTest{
		{"program main; 'Hello, ' + 'World!'", "Hello, World!"},
	}
	runVmTest(t, tests)
}

func TestBooleanLogic(t *testing.T) {
	tests := []vmTest{
		{"program main\ntrue", true},
		{"program main\nfalse", false},
		{"program main\n1.0 == 2.0", false},
		{"program main\n1.0 == 1.0", true},
		{"program main\ntrue == true", true},
		{"program main\nfalse == true", false},
		{"program main\n1.0 != 2.0", true},
		{"program main\n1.0 != 1.0", false},
		{"program main\ntrue != true", false},
		{"program main\nfalse != true", true},
		{"program main\n5 < 1", false},
		{"program main\n1 < 5", true},
		{"program main\n5 > 1", true},
		{"program main\n1 > 5", false},
		{"program main\n5 <= 1", false},
		{"program main\n1 <= 5", true},
		{"program main\n1 <= 1", true},
		{"program main\n5 >= 1", true},
		{"program main\n1 >= 5", false},
		{"program main\n5 >= 5", true},
		{"program main\ntrue || true", true},
		{"program main\nfalse || true", true},
		{"program main\nfalse || false", false},
		{"program main\ntrue && true", true},
		{"program main\nfalse && true", false},
		{"program main\nfalse && false", false},
		{"program main\n!false", true},
		{"program main\n!true", false},
	}

	runVmTest(t, tests)
}

func TestBlocks(t *testing.T) {
	tests := []vmTest{
		{"program main; do { 1 + 2 }", 3.0},
		{"program main; do ( 1 + 2 )", 3.0},
		{"program main; do { 1 + 1; 1 + 2 }", 3.0},
		{"program main; do ( 1 + 1; 1 + 2 )", 3.0},
		{"program main; do ( a = 10 ); a", 10},
	}

	runVmTest(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []vmTest{
		{`program main
		if true
			3.0
		else
			4.0`, 3.0},
		{"program main; if true; 3.0", 3.0},
		{"program main; if false 3.0", nil},
		{"program main; if false; 3.0; else; 4.0", 4.0},
	}

	runVmTest(t, tests)
}

func TestBindings(t *testing.T) {
	tests := []vmTest{
		{"program main; a = 10; a", 10.0},
		{"program main; a = 10; b = 10 * a", 100.0},
		{"program main; a = if true; 100", 100.0},
		//Lists are stored in the wrong order
		{"program main; a = [1, 2, 3, 4, 5, 6, 7, 8][3:6]; a", []float64{4, 5, 6, 7}},
		{"program main; a = 'Hello, World!'[3:7]; a", "lo, "},
		{"program main; a = [1, 2, 3, 4, 5, 6, 7][1]; a", 2},
		{"program main; a = 'Hello'[0]; a", "H"},
		{"program main; ten = func() do { a =10 }; ten()", 10},
	}

	runVmTest(t, tests)
}

func runVmTest(t *testing.T, tests []vmTest) {
	t.Helper()

	for _, test := range tests {
		tree := parse(test.input)

		comp := compiler.New()
		comp.Compile(tree)

		compiled := comp.GetCompiled()
		//fmt.Println(compiled) //Debug print statement, use it if you want

		vm := New(compiled)
		vm.Run()

		stackElement := vm.LastPopped()
		err := testExpectedVal(t, test.expected, stackElement)
		if err != nil {
			t.Errorf("%s", err)
		}

		if vm.stackPointer != 0 {
			t.Errorf("Stack not empty at test end %v ", vm.stack[:vm.stackPointer])
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

	if expected == nil {
		return testNilObj(got)
	}

	switch expected := expected.(type) {
	case float64:
		return testNumObj(float64(expected), got)
	case bool:
		return testBoolObj(bool(expected), got)
	case string:
		return testStrObj(string(expected), got)
	case []float64:
		return testListObj(t, []float64(expected), got)
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
		return fmt.Errorf("Expected Boolean Value.\nGot: %T (%+v)", got, got)
	}

	if result.Value != expected {
		return fmt.Errorf("Boolean Has Incorrect Value.\nGot: %v\nWant: %v", result.Value, expected)
	}

	return nil
}

func testNilObj(got values.Value) error {
	if got.Type() != values.NIL {
		return fmt.Errorf("Expected nil, Got: %T", got)
	}

	return nil
}

func testStrObj(expected string, got values.Value) error {
	result, ok := got.(values.String)
	if !ok {
		return fmt.Errorf("Expected string, got: %T", got)
	}

	if result.Value != expected {
		return fmt.Errorf("String has incorrect value\nGot: %T%+v\nWant: %v", result, result, expected)
	}

	return nil
}

func testListObj(t *testing.T, expected []float64, got values.Value) error {
	result, ok := got.(values.List)
	if !ok {
		return fmt.Errorf("Expected list, got %T", got)
	}

	for iter, i := range result.Values {
		err := testExpectedVal(t, expected[iter], i)
		if err != nil {
			return err
		}
	}

	return nil
}
