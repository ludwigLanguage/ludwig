package main

import (
	"fmt"
	"ludwig/src/lexer"
	"ludwig/src/parser"
	"ludwig/src/source"
	"os"
)

var version = "v0.1.5 -- Development"

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}

	switch os.Args[1] {
	case "-l":
		printToks()
	case "-p":
		printTree()
	case "-e":
		evalFile()
	case "-r":
		printHelp()
	default:
		printHelp()
	}
}

func printHelp() {
	msg := `
Welcome to the Ludwid Programming Language
Version Information: ` + version + `

    ludwig [option] | ludwig [option] [file]
    -h :: print this message and exit
    -l :: lex file, and print tokens
    -p :: parse file and print tree
    -e :: complete evaluation of the file
    `
	fmt.Println(msg)
	os.Exit(0)
}

func printToks() {
	if len(os.Args) < 3 {
		printHelp()
	}

	src := source.New(os.Args[2])
	lex := lexer.New(src)

	for !lex.IsDone() {
		fmt.Printf("%v -> %v\n", lex.CurTok.Alias, lex.CurTok)
		lex.MoveUp()
	}
}

func printTree() {
	if len(os.Args) < 3 {
		printHelp()
	}

	src := source.New(os.Args[2])
	lex := lexer.New(src)
	prs := parser.New(lex)
	prs.ParseProgram()

	prs.Tree.PrintAll("")
}

func evalFile() {
	printHelp()
}
