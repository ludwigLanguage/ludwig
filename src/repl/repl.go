package repl

import (
	"ludwig/src/evaluator"
	"ludwig/src/lexer"
	"ludwig/src/parser"
	"ludwig/src/source"
	"ludwig/src/values"

	"bufio"
	"fmt"
	"os"
)

const (
	PROMPT = "(ludwig) >> "
)

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)
	consts := values.NewSymTab()

	for {
		fmt.Print(PROMPT)
		text, _ := reader.ReadString('\n')

		src := source.NewWithStr(text, "repl")
		lex := lexer.New(src)
		prs := parser.New(lex)
		prs.ParseProgram()

		evaluator.EvalExpr(prs.Tree, consts)
	}
}
