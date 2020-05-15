/* This package creates a stream of single character strings out
 * of the contents of a given file name.
 */
package source

import (
	"io/ioutil"
	"ludwig/src/message"
	"os/user"
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

func New(inFile string) *Source {

	var contents string

	filename, _ := filepath.Abs(inFile)
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		contents = string(bytes)
	} else {

		usr, err := user.Current()
		if err != nil {
			message.Error(filename, "System", "Could not obtain user information", 0, 0)
		}

		sep := string(filepath.Separator)
		filename = usr.HomeDir + sep + ".ludwig" + sep + inFile
		bytes, err = ioutil.ReadFile(filename)

		if err != nil {
			message.Error(filename, "File", "Could not open this file", 0, 0)
		}

		contents = string(bytes)

	}

	return NewWithStr(contents, filename)
}

/* NewWithStr() must be public so that it can be used in the
 * evaluator when evaluating the "$" prefix
 */
func NewWithStr(contents string, filename string) *Source {
	s := &Source{}
	s.Filename = filename

	/* The extra spaces give curChar and nextChar
	 * something to be
	 */
	s.contents = contents + "  "

	s.curIter = -1
	s.nextIter = 0
	s.LineNo = 1

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
