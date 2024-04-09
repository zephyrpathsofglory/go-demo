package permute

/*
leetcode: 46给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

// 递归法，性能比较低：函数调用栈太深，栈的临时变量太多，导致内存占用过多
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

func permuteBacktrace(inputs []int) (res [][]int) {
	size := len(inputs)
	used := make([]bool, size)

	var dfs func(int, []int)
	dfs = func(depth int, path []int) {
		if depth == size {

			res = append(res, append([]int{}, path...))
			/* res = append(res, path)

			不能用上面的方式，因为 slice 底层是同一个数组，会随着后面 path 的修
			改而改变，比如下面的错误输出：
			[[5,4,2,6],[5,4,2,6],[5,6,2,4],[5,6,2,4],[5,2,6,4],[5,2,6,4],
			[4,5,2,6],[4,5,2,6],[4,6,2,5],[4,6,2,5],[4,2,6,5],[4,2,6,5],[6,5,2,4],
			[6,5,2,4],[6,4,2,5],[6,4,2,5],[6,2,4,5],[6,2,4,5],[2,5,6,4],[2,5,6,4],
			[2,4,6,5],[2,4,6,5],[2,6,4,5],[2,6,4,5]]
			为什么只有后 2位 重复？因为底层数组扩容，导致的底层数组地址发生改变，扩容顺序 1 -> 2 -> 4 。所以 3， 4 位置总是重复
			*/

			return
		}

		for i, j := range inputs {
			if used[i] {
				continue
			}

			path = append(path, j)
			used[i] = true
			dfs(depth+1, path)

			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})

	return
}
