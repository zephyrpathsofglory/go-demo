package od

import "sort"

/*
给定两个只包含数字的数组a，b，调整数组 a 里面数字的顺序，使得尽可能多的 a[i] > b[i]。数组 a
和 b 中的数字各不相同。输出所有可以达到最优结果的 a 数组的数量

输入描述
输入的第1行是数组 a 中的数字，其中只包含数字，每两个数字之间相隔一个空格，a 数组大小不超过 10
输入的第2行是数组 b 中的数字，其中只包含数字，每两个数字之间相隔一个空格，b 数组大小不超过 10
*/

func countSolutions(a, b []int) int {
	if len(b) > len(a) {
		b = b[:len(a)]
	}

	state := make([][]int, 0, len(b))

	for i := 0; i < len(b); i++ {
		state = append(state, make([]int, 0, len(a)))
	}

	sort.Ints(b)

	for i1, j := range b {
		for _, k := range a {
			if k > j {
				state[i1] = append(state[i1], k)
			}
		}

		if len(state[i1]) == 0 {
			break
		}
	}

	tt := make([]int, 0, len(state))
	for _, ss := range state {
		if len(ss) > 0 {
			tt = append(tt, len(ss))
		}
	}

	sort.Ints(tt)

	count := 0
	if len(tt) > 0 {
		count = 1
	}

	for idx, i := range tt {
		count *= (i - idx)
	}

	other := len(a) - len(tt)

	otherCount := 0
	if other > 0 {
		otherCount = 1
	}
	for other > 0 {
		otherCount *= other
		other -= 1
	}

	if otherCount == 0 {
		return count
	} else if count == 0 {
		return otherCount
	} else {
		return count * otherCount
	}
}

func contains(arr []int, d int) bool {
	for _, j := range arr {
		if d == j {
			return true
		}
	}

	return false
}

func remove(arr1, arr2 []int) (res []int) {
	for _, i := range arr1 {
		if contains(arr2, i) {
			continue
		}

		res = append(res, i)
	}

	return
}

func least(arr []int) int {
	l := arr[0]

	for _, i := range arr {
		if i < l {
			l = i
		}
	}

	return l
}
