package message

import (
	"log"
	"ludwig/src/tokens"
	"os"
	"strconv"
)

var (
	ShouldExit bool = true

	stderr = log.New(os.Stderr, "", 0)
)

func Error(f, n, m string, ln, cn int) {
	//Example: ./main.ldg (1:1) -- SyntaxError: No prefix parse func
	msg := f + " (" + strconv.Itoa(ln) + ":" + strconv.Itoa(cn) + ") -- " + n + "Error: " + m
	stderr.Println(msg)

	if ShouldExit {
		os.Exit(2)
	}
}

func RaiseError(n, m string, tok tokens.Token) {
	Error(tok.Filename, n, m, tok.LineNo, tok.ColumnNo)
}

func RuntimeErr(n, m string, tok tokens.Token, l *Log) {
	stderr.Println("Calls Made:")
	l.PrintStack()
	RaiseError(n, m, tok)
}

func VmError(n, m string, file string, lineNo int) {
	msg := file + " (" + strconv.Itoa(lineNo) + ") -- " + n + "Error: " + m
	stderr.Println(msg)

	os.Exit(2)
}

type Log struct {
	vals map[string]tokens.Token
}

func NewLog() *Log {
	return &Log{make(map[string]tokens.Token)}
}

func (l *Log) Add(fn string, tok tokens.Token) {
	l.vals[fn] = tok
}

func (l *Log) Rm(fn string) {
	delete(l.vals, fn)
}

func (l *Log) PrintStack() {
	for val, tok := range l.vals {

		stderr.Printf("\t%v (%v:%v) -- %v\n", tok.Filename, tok.LineNo, tok.ColumnNo, val)
	}
}
