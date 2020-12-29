package cerr

import "runtime/debug"

// golang error structure
// detailed logging than errors New struct

// New Create Error of error with method
func New(msg string) *cerr {
	return &cerr{s: msg}
}

// cerr BasicError structure
type cerr struct {
	s string
}

func (e *cerr) Error() string {
	debug.PrintStack()
	return e.s
}
