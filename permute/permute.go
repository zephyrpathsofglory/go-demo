package permute

/*
leetcode: 46
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

func permute(inputs []int) (res [][]int) {
	var f func([]int, []int)
	f = func(prev []int, nums []int) {
		if len(nums) == 1 {
			res = append(res, append(prev, nums[0]))
			return
		}

		for i, num := range nums {
			left := append([]int{}, nums[:i]...)
			right := append([]int{}, nums[i+1:]...)
			t := append(left, right...)
			f(append(prev, num), t)
			// f(append(prev, num), append(nums[0:i], nums[i+1:]...))
		}
	}

	f([]int{}, inputs)

	return
}
