package od

import (
	"cmp"
	"slices"
)

/*
你有 n 台机器编号为1-n，每台都需要完成一项工作， 机器经过配置后都能独立完成一项工作。 假设第i台机器你需要花 Bi 分钟进行
设置， 然后开始运行，Ji分钟后完成任务。 现在，你需要选择布置工作的顺序，使得用最短的时间完成所有工作。 注意，不能同时对
两台进行配置， 但配置完成的机器们可以同时执行他们各自的工作。

输入第一行输入代表总共有 M 组任务数据(1 < M <= 10); 每组数第一行为一个整数指定机器的数量N(0 < N <= 1000)。 随后的 N 行每
行两个整数，第一个表示B(0 <= B <= 10000), 第二个表示J(0 <= J <= 10000); 每组数据连续输入，不会用空行分割，各组任务单独计
时

输出描述对于每组任务，输出最短完成时间， 且每组的结果独占一行。 例如两组任务就应该有两行输出。
*/

// 运行时间长的先配置，因为越往后面，越晚结束，需要的时间越多
// 状态转移方程： dp[i] = max(dp[i-1], totalConfigurationTime+tasks[i][1])
// 网上有解答说是 dp[i] = Math.max(dp[i-1], dp[i-1] - machine[i-1][1] + machine[i][0] + machine[i][1])，这种是不对的，假设第一个
// 任务的时间足够大，加一个任务不会影响总的执行时间，如果是上面错误的方程，总的执行时间延长了，所以是错的

func taskScheduler(tasks [][2]int) int {
	slices.SortFunc(tasks, func(i, j [2]int) int {
		return cmp.Compare(j[1], i[1])
	})

	dp := make([]int, len(tasks), len(tasks))
	dp[0] = tasks[0][0] + tasks[0][1]

	totalConfigurationTime := tasks[0][0]
	for i := 1; i < len(tasks); i++ {
		totalConfigurationTime += tasks[i][0]
		dp[i] = max(dp[i-1], totalConfigurationTime+tasks[i][1])
	}

	return dp[len(tasks)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
