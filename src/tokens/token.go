/* If youu cant tell whats in this file, or don't know. 
 * Go away and stop messing with something you know nothing about
 * until you have spent some time learning. I recomend the book
 * "Writing an Interpreter in Go" By Thorsten Ball
 */
package tokens

type Token struct {
	Filename string
	LineNo   int
	ColumnNo int

	Value string
	Alias byte
}
