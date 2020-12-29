package cerr

import (
	"log"
	"testing"
)

func TestWriteOutErrorString(t *testing.T) {
	err := New(`original error`)
	log.Println(err)
}
