package message

import (
	"ludwig/src/tokens"
	"log"
	"os"
	"strconv"
)

var (
	shouldExit bool = true

	stderr = log.New(os.Stderr, "", 0)
)

func Error(f, n, m string, ln, cn int) {
	//Example: ./main.kgo (1:1) -- SyntaxError: No prefix parse func
	msg := f + " (" + strconv.Itoa(ln) + ":" + strconv.Itoa(cn) + ") -- " + n + "Error: " + m
	stderr.Println(msg)

	if shouldExit {
		os.Exit(2)
	}
}

func RaiseError(n, m string, tok tokens.Token) {
	Error(tok.Filename, n, m, tok.LineNo, tok.ColumnNo)
}
