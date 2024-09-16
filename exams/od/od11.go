package od

/*
题目描述
服务器连接方式包括直接相连，间接连接。
A 和 B 直接连接， B 和 C 直接连接，则 A 和 C 间接连接。直接连接和间接连接都可以发送广播。
给出一个 N * N 数组，代表 N 个服务器， matrix[i][j] == 1 ，则代表 i 和 j 直接连接；
不等于 1 时，代表 i 和 j 不直接连接。 matrix[i][i]== 1 ，即自己和自己直接连接。
matrix[i][j]==matrix[j][i] 。
计算初始需要给几台服务器广播，才可以使侮个服务器都收到广播。
*/

/*
实质是图的遍历，求连通分量的个数；
*/

func computeBroadcastCount(connectivities [][]int) int {
	traveled := make([]bool, len(connectivities))
	count := 0

	var f func(int)
	f = func(n int) {
		if traveled[n] {
			return
		}

		traveled[n] = true

		for j, b := range connectivities[n] {
			if b == 1 && j != n {
				f(j)
			}
		}
	}

	for i := range connectivities {
		if !traveled[i] {
			f(i)
			count += 1
		}
	}

	return count
}
