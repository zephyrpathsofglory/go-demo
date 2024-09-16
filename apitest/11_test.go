package apitest

import (
	"fmt"
	"testing"
)

/*
[0 2]
[1]
[9 3 4 5]
[9 3 4]
*/

func Test11(_ *testing.T) {
	arr := []int{1}

	myfunc1(arr)
	fmt.Println(arr)

	arr = append(arr, 3)
	arr = append(arr, 4)
	myfunc2(arr)
	fmt.Println(arr)
}

func myfunc1(arr []int) {
	arr = append(arr, 2)
	arr[0] = 0
	fmt.Println(arr)
}

func myfunc2(arr []int) {
	arr = append(arr, 5)
	arr[0] = 9
	fmt.Println(arr)
}
