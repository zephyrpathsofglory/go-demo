/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

输入： heights = [2,4]
输出： 4
*/
package leetcode

func LargestRectangleArea(heights []int) int {
	max := 0
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)

	stack := make([]int, 0, len(heights))

	stack = append(stack, 0)

	for i := 1; i < len(heights); i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prev := stack[len(stack)-1]

			size := heights[last] * (i - prev - 1)

			if size > max {
				max = size
			}
		}

		stack = append(stack, i)
	}

	return max
}
