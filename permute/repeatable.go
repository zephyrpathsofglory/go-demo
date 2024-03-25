package permute

/*
leetcode 47
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/

func repeatable(inputs []int) (res [][]int) {
	var f func([]int, []int)

	f = func(preset []int, left []int) {
		if len(left) == 1 {
			res = append(res, append(preset, left[0]))
			return
		}

		appearance := make(map[int]bool)
		for i, n := range left {
			if !appearance[n] {
				t1 := append([]int{}, left[0:i]...)
				t2 := append(t1, left[i+1:]...)
				f(append(preset, n), t2)
				appearance[n] = true
			}
		}
	}

	f([]int{}, inputs)

	return
}
