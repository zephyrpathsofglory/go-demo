package knapsack

/*
订单列表里面，每个单可能有不同的乘客数 1，2，3个， 车总共可以坐 4人，
优先选择前面的订单，确保尽可能不空座位,返回选择的订单的idx数组
*/

func totalSum(nums []int, total int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	state := make([][][]int, len(nums))
	for i := range state {
		state[i] = make([][]int, total)
		for j := range state[i] {
			state[i][j] = []int{}
		}
	}

	f := func(i, j int) ([]int, int) {
		if i < 0 || j < 0 {
			return []int{}, 0
		} else {
			s := state[i][j]
			c := 0
			for _, i := range s {
				c += nums[i]
			}
			return s, c
		}
	}

	for i, num := range nums {
		for j := 0; j < total; j++ {
			arr1, c1 := f(i-1, j)
			arr2, c2 := f(i-1, j-num)

			c := c2 + num

			if c == total {
				return append(arr2, i)
			}

			if c > c1 && c <= j+1 {
				state[i][j] = append(append([]int{}, arr2...), i)
			} else {
				state[i][j] = append([]int{}, arr1...)
			}
		}
	}

	return state[len(nums)-1][total-1]
}
