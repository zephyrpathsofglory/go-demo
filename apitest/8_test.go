package apitest

import (
	"fmt"
	"testing"
)

func Test8(_ *testing.T) {
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	quarter := months[3:6]
	fmt.Printf("%v, %d,%d", quarter, len(quarter), cap(quarter))

	extended := quarter[:4]
	fmt.Printf("%v, %d, %d", extended, len(extended), cap(extended))
}
