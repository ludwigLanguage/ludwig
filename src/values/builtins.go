package values

import (
	"bufio"
	"fmt"
	"ludwig/src/message"
	"ludwig/src/tokens"
	"os"
	"os/exec"
	"strconv"
)

/////////////////////////////////////////////////

type Builtin struct {
	Fn func([]Value, tokens.Token, *message.Log) Value
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

func print(v []Value, tok tokens.Token, l *message.Log) Value {
	if !(len(v) >= 1) {
		message.RuntimeErr("Argument", "Must have at least one argument to 'print'", tok, l)
	}

	var rtrn string
	for j, i := range v {
		rtrn += i.Stringify()

		if j != len(v)-1 {
			rtrn += " "
		}
	}

	fmt.Print(rtrn)
	return &String{rtrn, tok}
}

////////////////////////////////////
func println(v []Value, tok tokens.Token, l *message.Log) Value {
	values := append(v, &String{"\n", tok})
	return print(values, tok, l)
}

///////////////////////////////////
func read(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 2 {
		message.RuntimeErr("Argument", "read() must have two arguments", tok, l)
	}

	if len(v[1].Stringify()) != 1 {
		message.RuntimeErr("Argument", "read()'s second argument must be 1 character in length", tok, l)
	}

	print([]Value{v[0]}, tok, l)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString(v[1].Stringify()[0])

	if err != nil {
		message.RuntimeErr("Argument", "failed to read input", tok, l)
	}

	return &String{text, tok}
}

//////////////////////////////////
func typeOf(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 1 {
		message.RuntimeErr("Argument", "type_of() Must have exactly one argument", tok, l)
	}
	t := &TypeIdent{v[0].Type(), tok}
	return t
}

/////////////////////////////////////////////////

func str(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 1 {
		message.RuntimeErr("Argument", "str() must have exactly one argument", tok, l)
	}

	return &String{v[0].Stringify(), tok}
}

/////////////////////////////////////////////////

func num(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 1 {
		message.RuntimeErr("Argument", "num() must have exactly one argument", tok, l)
	} else if v[0].Type() != STR {
		message.RuntimeErr("Type", "num() must have a string as the argument", tok, l)
	}

	flt, err := strconv.ParseFloat(v[0].(*String).Value, 64)

	if err != nil {
		message.RuntimeErr("Type", "Cannot convert this into a number", tok, l)
	}

	return &Number{flt, tok}
}

/////////////////////////////////////////////////

func Length(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 1 {
		message.RuntimeErr("Argument", "'len' must have one argument", tok, l)
	}

	switch val := v[0].(type) {
	case *String:
		return &Number{float64(len(val.Value)), val.Tok}
	case *List:
		return &Number{float64(len(val.Values)), val.Tok}
	default:
		message.RuntimeErr("Type", "Expected list or string on 'len' call", val.GetTok(), l)
	}

	return &Nil{tok}
}

/////////////////////////////////////////////////
func osCall(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) < 2 {
		message.RuntimeErr("Argument", "'system' must have two arguments", tok, l)
	}

	if v[0].Type() != BOOL {
		message.RuntimeErr("Type", "First argument of 'system' must be a boolean", v[0].GetTok(), l)
	}
	shouldDisplayOutput := v[0].(*Boolean).Value

	var commandName string
	var commandArgs = []string{}
	for j, i := range v[1:] {
		if i.Type() != STR {
			message.RuntimeErr("Type", "Expected argument type 'String'", v[0].GetTok(), l)
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
	objSymTab.SetVal("output", &String{out, v[0].GetTok()})
	objSymTab.SetVal("error", &String{err, v[0].GetTok()})

	if shouldDisplayOutput {
		fmt.Print(out, err)
	}

	return &Object{objSymTab, v[0].GetTok()}
}

/////////////////////////////////////////////////

func osExit(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 1 {
		message.RuntimeErr("Argument", "'exit' must have exactly one argument", tok, l)
	}

	if v[0].Type() != "Number" {
		message.RuntimeErr("Type", "First argument of 'exit' must be a number", v[0].GetTok(), l)
	}
	var exitCode int = int(v[0].(*Number).Value)
	os.Exit(exitCode)
	return nil

}

///////////////////////////////////////////////

func panic(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 2 {
		message.RuntimeErr("Argument", "Panic must have 2 arguments", tok, l)
	}

	message.RuntimeErr(v[0].Stringify(), v[1].Stringify(), tok, l)
	return v[0]
}

////////////////////////////////////////////////

func check_type(v []Value, tok tokens.Token, l *message.Log) Value {
	if len(v) != 2 {
		message.RuntimeErr("Argument", "check_type must have 2 arguments", tok, l)
	}

	if v[0].Type() != TYPE_IDENT {
		message.RuntimeErr("Type", "expected a type identifier as the first argument", tok, l)
	}

	expect := v[0].(*TypeIdent).Value
	if expect != v[1].Type() {
		message.RuntimeErr("Type", "Expected type '"+expect+"' got '"+v[1].Type()+"'", v[1].GetTok(), l)
	}

	return &Nil{tok}
}

///////////////////////////////////////////////
var BuiltinsMap = map[string]Value{
	"print":      &Builtin{print},
	"println":    &Builtin{println},
	"read":       &Builtin{read},
	"type_of":    &Builtin{typeOf},
	"str":        &Builtin{str},
	"num":        &Builtin{num},
	"len":        &Builtin{Length},
	"system":     &Builtin{osCall},
	"exit":       &Builtin{osExit},
	"panic":      &Builtin{panic},
	"check_type": &Builtin{check_type},
}
