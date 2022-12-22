package c29_v2

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Interface(t *testing.T) {
	err := returnsError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}
