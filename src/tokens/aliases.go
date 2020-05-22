/* This file holds the various token aliases as constants.
 */
package tokens

const (
	EOF byte = iota
	EOL

	IDENT
	NUM
	STR
	BOOL
	NIL

	LBRACK
	RBRACK
	LCURL
	RCURL
	LPAREN
	RPAREN

	POP

	OP1
	OP2
	OP3
	OP4
	OP5

	DOT

	COLON
	COMMA

	IF
	ELSE
	FN
	IMPORT
	FOR
	IN
	WHILE
	PROG
	PACK
	PRIV
	PUB
	CLASS

	TYPE_INDENT
	BUILTIN
)
