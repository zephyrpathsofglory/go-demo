package apitest

import (
	"errors"
	"fmt"
	"testing"
)

/*
nil
c
a
*/

func f() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()

	defer func(err error) {
		fmt.Println(err)
		err = errors.New("b")
	}(err)

	err = errors.New("c")

	return
}

func Test10(_ *testing.T) {
	fmt.Println(f())
}
