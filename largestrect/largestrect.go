package largestrect

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

			left := stack[len(stack)-1]

			size := heights[last] * (i - left - 1)

			if size > max {
				max = size
			}
		}

		stack = append(stack, i)
	}

	return max
}
