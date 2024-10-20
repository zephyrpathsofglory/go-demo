/*
专业级第二题：最少committer

一个 MR 修改了 files， files[i] 表示允许评审这个文件的 committer 编号列表，每个文件需要至少一个 committer 评审才能合并，每名 committer
可以评审多个文件

请计算并返回该 MR 最少需要多少名 committer 才能完成所有文件的评审？
*/
package od

func LeastCommitter(files [][]int) int {
	var dfs func([][]int) int
	dfs = func(fileCommitter [][]int) int {
		minCommitter := len(fileCommitter)
		if len(fileCommitter) == 0 {
			return 0
		}

		for _, committer := range fileCommitter[0] {
			left := [][]int{}
			for _, commiters := range fileCommitter {
				if !contains(commiters, committer) {
					left = append(left, commiters)
				}
			}
			count := 1 + dfs(left)
			if count < minCommitter {
				minCommitter = count
			}
		}

		return minCommitter
	}

	return dfs(files)
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}
