package values

import (
	"ludwig/src/message"
	"ludwig/src/tokens"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

/////////////////////////////////////////////////

type Builtin struct {
	Fn func([]Value, tokens.Token) Value
}

func (b *Builtin) GetTok() tokens.Token {
	return tokens.Token{"ludwig/src/evaluator", 0, 0, "", tokens.RBRACK}
}

func (b *Builtin) Stringify() string {
	return "builtinFn()"
}

func (b *Builtin) Type() string {
	return BUILTIN
}

////////////////////////////////////////////

func print(v []Value, tok tokens.Token) Value {
	if !(len(v) >= 1) {
		message.RaiseError("Argument", "Must have at least one argument to 'print'", tok)
	} 

	var rtrn string
	for j, i := range v {
		rtrn += i.Stringify()
		fmt.Print(i.Stringify())

		if j != len(v)-1 {
			rtrn += " "
			fmt.Print(" ")
		}
	}

	return &String {rtrn, tok}
}

////////////////////////////////////
func println(v []Value, tok tokens.Token) Value {
	values := append(v, &String {"\n", tok})
	return print(values, tok)
}
///////////////////////////////////

func typeOf(v []Value, tok tokens.Token) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "typeOf() Must have exactly one argument", tok)
	}
	return &String{v[0].Type(), tok}
}

/////////////////////////////////////////////////

func str(v []Value, tok tokens.Token) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "str() must have exactly one argument", tok)
	}

	return &String{v[0].Stringify(), tok}
}

/////////////////////////////////////////////////

func num(v []Value, tok tokens.Token) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "num() must have exactly one argument", tok)
	} else if v[0].Type() != STR {
		message.RaiseError("Type", "num() must have a string as the argument", tok)
	}

	flt, err := strconv.ParseFloat(v[0].(*String).Value, 64)

	if err != nil {
		message.RaiseError("Type", "Cannot convert this into a number", tok)
	}

	return &Number{flt, tok}
}

/////////////////////////////////////////////////

func length(v []Value, tok tokens.Token) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "'len' must have one argument", tok)
	}

	switch val := v[0].(type) {
	case *String:
		return &Number{float64(len(val.Value)), val.Tok}
	case *List:
		return &Number{float64(len(val.Values)), val.Tok}
	default:
		message.RaiseError("Type", "Expected list or string on 'len' call", val.GetTok())
	}

	return &Nil {tok}
}

/////////////////////////////////////////////////
func osCall(v []Value, tok tokens.Token) Value {
	if len(v) < 2 {
		message.RaiseError("Argument", "'system' must have two arguments", tok)
	}
	
	typeOfArg1 := fmt.Sprintf("%T", v[0])
	if typeOfArg1 != "*values.Boolean" {
		message.RaiseError("Type", "First argument of 'system' must be a boolean", v[0].GetTok())
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

	rawOut, rawErr := cmd.CombinedOutput()

	var out, err string
	out = string(rawOut)

	if rawErr == nil {
		err = ""
	} else {
		err = rawErr.Error()
	}

	objSymTab := NewSymTab()
	objSymTab.SetVal("output", &String { out, v[0].GetTok() })
	objSymTab.SetVal("error", &String { err, v[0].GetTok() })

	if shouldDisplayOutput {
		fmt.Print(out, err)
	}

	return &Object {objSymTab, v[0].GetTok()}
}

/////////////////////////////////////////////////

func osExit(v []Value, tok tokens.Token) Value {
	if len(v) != 1 {
		message.RaiseError("Argument", "'exit' must have exactly one argument", tok)
	}
	
	if v[0].Type() != "Number" {
		message.RaiseError("Type", "First argument of 'exit' must be a number", v[0].GetTok())
	}
	var exitCode int = int(v[0].(*Number).Value)
	os.Exit(exitCode)
	return nil

}

///////////////////////////////////////////////

///////////////////////////////////////////////
var BuiltinsMap = map[string]Value { 
	"print":   &Builtin{print},
	"println": &Builtin{println},
	"typeOf":  &Builtin{typeOf},
	"str":     &Builtin{str},
	"num":     &Builtin{num},
	"len":     &Builtin{length},
	"system":  &Builtin{osCall},
	"exit":	   &Builtin{osExit},
}
