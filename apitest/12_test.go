package apitest

import (
	"fmt"
	"testing"
)

/*
0
0
3
3
0
4
0
5
*/

func Test12(_ *testing.T) {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
