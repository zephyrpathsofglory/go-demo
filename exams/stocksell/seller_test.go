package stocksell_test

import (
	"testing"

	"github.com/go-demo/exams/stocksell"
	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	{
		prices := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 4, stocksell.MaxProfit(prices))
	}
	{
		prices := []int{1, 20, 2, 3, 4}
		assert.Equal(t, 19, stocksell.MaxProfit(prices))
	}
	{
		prices := []int{1, 5, 2, 7, 3, 4, 6, 2}
		assert.Equal(t, 6, stocksell.MaxProfit(prices))
	}
	{
		prices := []int{10, 20, 1, 2, 9}
		assert.Equal(t, 10, stocksell.MaxProfit(prices))
	}
	{
		prices := []int{5, 4, 3, 2, 1}
		assert.Equal(t, 0, stocksell.MaxProfit(prices))
	}
	{
		prices := []int{7, 1, 5, 3, 6, 4}
		assert.Equal(t, 5, stocksell.MaxProfit(prices))
	}
}
