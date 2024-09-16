package apitest

import (
	"fmt"
	"testing"
)

func MyAppend(arr []int) []int {
	arr = append(arr, 5)

	return arr
}

func MyAdd(arr []int) {
	for i := range arr {
		arr[i] = arr[i] + 5
	}

	return
}

func Test9(_ *testing.T) {
	arr := []int{1, 2, 3, 4}
	MyAppend(arr)
	fmt.Println(arr)

	MyAdd(arr)
	fmt.Println(arr)
}
