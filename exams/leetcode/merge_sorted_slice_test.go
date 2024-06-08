package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedSlice(t *testing.T) {
	type test struct {
		nums1 []int
		nums2 []int
		want  []int
	}

	tests := []test{
		{
			[]int{1, 2, 3},
			[]int{2, 5, 6},
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			[]int{1},
			[]int{2, 3},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		res := mergeSortedSlice(tt.nums1, tt.nums2)
		assert.Equal(t, tt.want, res)
	}
}
