package selection_test

import (
	"testing"

	"github.com/go-demo/sort/selection"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	{
		arr := []int{1}
		selection.SelectionSort(arr)
		assert.Equal(t, []int{1}, arr)
	}

	{
		arr := []int{2, 3, 1}
		selection.SelectionSort(arr)
		assert.Equal(t, []int{1, 2, 3}, arr)
	}

	{
		arr := []int{2, 1}
		selection.SelectionSort(arr)
		assert.Equal(t, []int{1, 2}, arr)
	}

	{
		arr := []int{5, 27, 2, 673, 73}
		selection.SelectionSort(arr)
		assert.Equal(t, []int{2, 5, 27, 73, 673}, arr)
	}
}
