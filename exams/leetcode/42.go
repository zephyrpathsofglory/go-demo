/* 42. Trapping Rain Water

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

输入：height = [4,2,0,3,2,5]
输出：9

*/

package leetcode

type maxValue struct {
	idxVal        int
	leftMaxValue  int
	rightMaxValue int
}

func TrappingRainWater(nums []int) int {
	mvs := make([]*maxValue, len(nums))

	max := 0
	for i, num := range nums {
		mv := maxValue{
			idxVal: num,
		}

		if num >= max {
			max = num
			mv.leftMaxValue = num
		} else {
			mv.leftMaxValue = max
		}

		mvs[i] = &mv
	}

	max = 0
	for i := len(nums) - 1; i >= 0; i-- {
		mv := mvs[i]
		num := nums[i]
		if num >= max {
			max = num
			mv.rightMaxValue = num
		} else {
			mv.rightMaxValue = max
		}
	}

	total := 0
	for _, mv := range mvs {
		total += (minV(mv.leftMaxValue, mv.rightMaxValue) - mv.idxVal)
	}

	return total
}

func minV(x, y int) int {
	if x < y {
		return x
	}

	return y
}
