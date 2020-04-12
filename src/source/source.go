/* This package creates a stream of single character strings out
 * of the contents of a given file name. 
 */ 
package source

import (
	"ludwig/src/message"
	"io/ioutil"
	"path/filepath"
)

var (
	EOF byte = 0
)

type Source struct {
	Filename string
	contents string

	LineNo   int
	ColumnNo int

	curIter  int
	nextIter int

	CurChar  byte
	NextChar byte
}

func New(filename string) *Source {

	var contents string

	filename, _ = filepath.Abs(filename)
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		contents = string(bytes)
	} else {
		message.Error(filename, "File", "Could not open this file", 0, 0)
	}

	return NewWithStr(contents, filename)
}

/* NewWithStr() must be public so that it can be used in the
 * evaluator when evaluating the "$" prefix
 */
func NewWithStr(contents string, filename string) *Source {
	s := &Source {}
	s.Filename = filename

	/* The addition of the do...end block serves two purposes
	 * 1) It allows the parser to process the file as a 
	 *    block of expressions so that we dont have to
	 *    create any speacial protocol to process a program
	 * 2) If the file is empty, it creates two characters to
	 *    become s.CurChar, and s.NextChar so that we do not
	 *    fail out when we try to assign those
	 */
	s.contents = "do\n" + contents + "\nend"

	s.curIter = -1
	s.nextIter = 0

	s.MoveUp()

	return s
}

func (s *Source) MoveUp() {
	s.moveItersUp()
	s.setLocation()
	length := len(s.contents)

	if length < s.nextIter {
		s.CurChar = EOF
		s.NextChar = EOF

	} else if length == s.nextIter {
		s.NextChar = EOF
		s.CurChar = s.contents[s.curIter]

	} else {
		s.NextChar = s.contents[s.nextIter]
		s.CurChar = s.contents[s.curIter]

	}

}

func (s *Source) IsDone() bool {
	return s.CurChar == EOF
}

func (s *Source) moveItersUp() {
	s.curIter++
	s.nextIter++
}

func (s *Source) setLocation() {
	if s.CurChar == '\n' {
		s.ColumnNo = 0
		s.LineNo += 1
	}
	s.ColumnNo++
}
