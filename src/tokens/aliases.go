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
	STRUCT
	IMPORT
	DO
	END
	FOR
	IN
	WHILE

	TYPE_INDENT
)
