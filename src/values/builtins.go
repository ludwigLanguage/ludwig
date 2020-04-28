package values

import (
	"bufio"
	"fmt"
	"ludwig/src/ast"
	"ludwig/src/message"
	"ludwig/src/tokens"
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

func (b *Builtin) ConvertToAst() ast.Node {
	return &ast.Nil{tokens.Token{"ludwig/src/evaluator", 0, 0, "", tokens.RBRACK}}
}

////////////////////////////////////////////

func print(v []Value, tok tokens.Token) Value {
	if !(len(v) >= 1) {
		message.RaiseError("Argument", "Must have at least one argument to 'print'", tok)
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
func println(v []Value, tok tokens.Token) Value {
	values := append(v, &String{"\n", tok})
	return print(values, tok)
}

///////////////////////////////////
func read(v []Value, tok tokens.Token) Value {
	if len(v) != 2 {
		message.RaiseError("Argument", "read() must have two arguments", tok)
	}

	if len(v[1].Stringify()) != 1 {
		message.RaiseError("Argument", "read()'s second argument must be 1 character in length", tok)
	}

	print([]Value{v[0]}, tok)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString(v[1].Stringify()[0])

	if err != nil {
		message.RaiseError("Argument", "failed to read input", tok)
	}

	return &String{text, tok}
}

//////////////////////////////////
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

func Length(v []Value, tok tokens.Token) Value {
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

	return &Nil{tok}
}

/////////////////////////////////////////////////
func osCall(v []Value, tok tokens.Token) Value {
	if len(v) < 2 {
		message.RaiseError("Argument", "'system' must have two arguments", tok)
	}

	if v[0].Type() != ast.BOOL {
		message.RaiseError("Type", "First argument of 'system' must be a boolean", v[0].GetTok())
	}
	shouldDisplayOutput := v[0].(*Boolean).Value

	var commandName string
	var commandArgs = []string{}
	for j, i := range v[1:] {
		if i.Type() != STR {
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
	objSymTab.SetVal("output", &String{out, v[0].GetTok()})
	objSymTab.SetVal("error", &String{err, v[0].GetTok()})

	if shouldDisplayOutput {
		fmt.Print(out, err)
	}

	return &Object{objSymTab, v[0].GetTok()}
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

func panic(v []Value, tok tokens.Token) Value {
	if len(v) != 2 {
		message.RaiseError("Argument", "Panic must have 2 arguments", tok)
	}

	message.RaiseError(v[0].Stringify(), v[1].Stringify(), tok)
	return v[0]
}

///////////////////////////////////////////////
var BuiltinsMap = map[string]Value{
	"print":   &Builtin{print},
	"println": &Builtin{println},
	"read":    &Builtin{read},
	"typeOf":  &Builtin{typeOf},
	"str":     &Builtin{str},
	"num":     &Builtin{num},
	"len":     &Builtin{Length},
	"system":  &Builtin{osCall},
	"exit":    &Builtin{osExit},
	"panic":   &Builtin{panic},
}
