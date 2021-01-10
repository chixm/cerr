package cerr

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
)

// golang error structure
// detailed logging than errors New struct

// New Create Error of error with method
func New(msg string) error {
	return &cerr{s: fmt.Sprintf(caller(2), msg)}
}

// Wrap wraps the error and add new line with error files
func Wrap(err error, msg string) error {
	e := &cerr{s: fmt.Sprintf(caller(3), msg)}
	c := &cerr{}
	ok := errors.As(e, &c)
	if !ok {
		log.Panicln(`failed to wrap error`)
		return err
	}
	c.stack = err
	return c
}

// cerr BasicError structure
type cerr struct {
	s     string
	stack error
}

func (e *cerr) Error() string {
	if e.stack == nil {
		return e.s
	} else {
		return fmt.Sprintf(`%s <[%s]`, e.s, e.stack.Error())
	}
}

func caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		f := filepath.Base(file)
		return f + `[` + strconv.Itoa(line) + `] %s`
	}
	return `%s`
}
