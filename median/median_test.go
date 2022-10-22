package median_test

import (
	"testing"

	"github.com/go-demo/median"
	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	{
		nums1 := []int{1, 2, 3, 4}
		nums2 := []int{5, 6, 7, 8}
		assert.Equal(t, 4.5, median.Median(nums1, nums2))
	}

	{
		nums1 := []int{}
		nums2 := []int{1}
		assert.Equal(t, 1.0, median.Median(nums1, nums2))
	}

	{
		nums1 := []int{}
		nums2 := []int{2, 3}
		assert.Equal(t, 2.5, median.Median(nums1, nums2))
	}

	{
		nums1 := []int{10}
		nums2 := []int{1, 2, 3}
		assert.Equal(t, 2.5, median.Median(nums1, nums2))
	}

	{
		nums1 := []int{2, 4, 7, 9}
		nums2 := []int{1, 5, 6, 10}
		assert.Equal(t, 5.5, median.Median(nums1, nums2))
	}

	{
		nums1 := []int{2, 4, 7, 9}
		nums2 := []int{1, 5, 6, 8, 11}
		assert.Equal(t, 6.0, median.Median(nums1, nums2))
	}
}
