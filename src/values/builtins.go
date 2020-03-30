package values

import (
	"ludwig/src/message"
	"ludwig/src/tokens"

	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var (
	TOK       = tokens.Token{"ludwig/src/evaluator", 0, 0, "", tokens.RBRACK}
	NILRTRN Value = &Nil{ TOK}
)

/////////////////////////////////////////////////

type Builtin struct {
	Fn func([]Value) Value
}

func (b *Builtin) GetTok() tokens.Token {
	return TOK
}

func (b *Builtin) Stringify() string {
	return "builtinFn()"
}

func (b *Builtin) Type() string {
	return BUILTIN
}

/////////////////////////////////////////////////

func write(v []Value) Value {
	for _, i := range v {
		fmt.Printf("%v", i.Stringify())
	}

	return NILRTRN
}

/////////////////////////////////////////////////

func writeln(v []Value) Value {
	write(v)
	fmt.Println() //For the newline

	return NILRTRN
}

/////////////////////////////////////////////////

func readln(v []Value) Value {
	write(v)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return &String{text[:len(text)-1], TOK} //Cut '\n' off
}

/////////////////////////////////////////////////

func typeOf(v []Value) Value {
	if len(v) != 1 {
		message.Error(
			"Unknown.ludwig",
			"Builtin",
			"'typeOf' must have 1 argument",
			0, 0)
	}

	return &String{v[0].Type(), v[0].GetTok()}
}

/////////////////////////////////////////////////

func str(v []Value) Value {
	if len(v) != 1 {
		message.Error(
			"Unknown.ludwig",
			"Builtin",
			"'str' must have 1 argument",
			0, 0)
	}

	return &String{v[0].Stringify(), TOK}
}

/////////////////////////////////////////////////

func num(v []Value) Value {
	if len(v) != 1 {
		message.Error(
			"Unknown.ludwig",
			"Builtin",
			"'num' must have 1 argument",
			0, 0)
	} else if v[0].Type() != STR {
		message.Error(
			"Unknown.ludwig",
			"Builtin",
			"'num' argument must be a string",
			0, 0)
	}

	flt, err := strconv.ParseFloat(v[0].(*String).Value, 64)

	if err != nil {
		message.RaiseError("Type", "Cannot convert this into a number", v[0].GetTok())
	}

	return &Number{flt, v[0].GetTok()}
}

/////////////////////////////////////////////////

func length(v []Value) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "'len' must have one argument", TOK)
	}

	switch val := v[0].(type) {
	case *String:
		return &Number{float64(len(val.Value)), val.Tok}
	case *List:
		return &Number{float64(len(val.Values)), val.Tok}
	default:
		message.RaiseError("Type", "Expected list or string on 'len' call", val.GetTok())
	}

	return NILRTRN
}

/////////////////////////////////////////////////
func osCall(v []Value) Value {
	if len(v) < 2 {
		message.RaiseError("Argument", "'os' must have two arguments", TOK)
	}
	
	typeOfArg1 := fmt.Sprintf("%T", v[0])
	if typeOfArg1 != "*values.Boolean" {
		message.RaiseError("Type", "First argument of 'os' must be a boolean", v[0].GetTok())
	}
	shouldDisplayOutput := v[0].(*Boolean).Value

	var commandName string
	var commandArgs = []string{}
	for j, i := range v[1:] {
		typeOfArg := fmt.Sprintf("%T", i)
		if typeOfArg != "*values.String" {
			message.RaiseError("Type", "Expected argument type 'String'", v[0].GetTok())
		}

		if j == 0 {
			commandName = i.(*String).Value
		} else {
			commandArgs = append(commandArgs, i.(*String).Value)
		}
	}
	cmd := exec.Command(commandName, commandArgs...)
	output, err := cmd.CombinedOutput()

	var cmdErrVal Value
	if err == nil {
		cmdErrVal = &Nil {v[0].GetTok()}
	} else {
		cmdErrVal = &String {err.Error(), v[0].GetTok()}
	}

	cmdOutputVal := &String {string(output), v[0].GetTok()}

	if shouldDisplayOutput {
		fmt.Printf(string(output))
	}

	objSymTab := NewSymTab()
	objSymTab.SetVal("Output", cmdOutputVal)
	objSymTab.SetVal("Error", cmdErrVal)

	return &Object {objSymTab, v[0].GetTok()}
}

/////////////////////////////////////////////////

func osExit(v []Value) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "'exit' must have only one argument", TOK)
	}
	
	typeOfArg1 := fmt.Sprintf("%T", v[0])
	if typeOfArg1 != "*values.Number" {
		message.RaiseError("Type", "First argument of 'exit' must be a number", v[0].GetTok())
	}
	var exitCode int = int(v[0].(*Number).Value)
	os.Exit(exitCode)
	return nil

}

////////////////////////////////////////////////

var BuiltinsMap = map[string]Value{
	"println": &Builtin{writeln},
	"print":   &Builtin{write},
	"read":    &Builtin{readln},
	"typeOf":  &Builtin{typeOf},
	"str":     &Builtin{str},
	"num":     &Builtin{num},
	"len":     &Builtin{length},
	"system":  &Builtin{osCall},
	"exit":	   &Builtin{osExit},
}
