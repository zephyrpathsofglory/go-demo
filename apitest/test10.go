package apitest

import (
	"errors"
	"fmt"
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

func main10() {
	fmt.Println(f())
}
