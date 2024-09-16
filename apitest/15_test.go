package apitest

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3}
	p1 := getArrayPointer(a)

	b := []int{4}
	c := []int{4}
	b = append(b, a[len(a)-2], a[len(a)-1])
	c = append(c, a[len(a)-2:]...)

	a = append(a, 6, 7, 8, 9, 10, 11, 12, 13)
	p2 := getArrayPointer(a)

	assert.NotEqual(t, p1, p2)

	a[1] = 11
	assert.Equal(t, []int{4, 2, 3}, b)
	assert.Equal(t, []int{4, 2, 3}, c)
}

func getArrayPointer(slice interface{}) uintptr {
	// 使用反射获取切片的值
	val := reflect.ValueOf(slice)
	// 确保传入的是切片类型
	if val.Kind() != reflect.Slice {
		panic("expected a slice")
	}
	// 获取底层数组的指针
	return val.Pointer()
}
