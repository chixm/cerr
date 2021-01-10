package cerr

import (
	"errors"
	"log"
	"testing"
)

func TestWriteOutErrorString(t *testing.T) {
	err := New(`original error`)
	log.Println(err)
	// log should write this file name. stack trace shows the line of New method is located
}

func TestWrapErrorOfOtherPackage(t *testing.T) {
	ane := errors.New(`another error`)
	err := Wrap(ane, `layered errors`)
	log.Println(err)
}
