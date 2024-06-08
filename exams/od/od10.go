package od

import "slices"

/*
题目描述

火锅里会在不同时间下很多菜. 不同食材要煮不同的时间，才能变得刚好合适。 你希望吃到最多的刚好合适的菜，
但你的手速不够快，用 m 代表手速，每次下手捞菜后至少要过 m 秒才能再捞(每次只能捞一个)。 那么用最合理的策略，最多能吃到多少刚好合适的菜?

输入描述

第一行两个整数 n，m，其中 n 代表往锅里下的菜的个数，m 代表手速。 (1<n，m< 1000)

接下来有 n 行，每行有两个数 x，y 代表第 x 秒下的菜过 y 秒才能变得刚好合适。

(1<x,y< 1000)

输出描述

输出一个整数代表用最合理的策略，最多能吃到刚好合适的菜的数量.

注意： 菜是在不同时间下的，即每个菜的下菜时间不相同
*/

// 贪心算法，能拿就拿， 假设 第一，二个菜时间点分别为 A,B, 如果先拿B，下个可以拿的菜为 C， 那么先拿A，也可以拿 C， 相反，如果先拿 A，下一个可拿
// 的为 D，先拿B不一定能拿 D。所以先拿总是多拿

func chiHuoGuo(xiaCai [][2]int, handSpeed int) int {
	cooked := make([]int, 0, len(xiaCai))

	for _, cai := range xiaCai {
		cooked = append(cooked, cai[0]+cai[1])
	}

	slices.Sort(cooked)

	count := 1

	nextCai := cooked[0] + handSpeed

	for i := 1; i < len(cooked); i++ {
		if cooked[i] >= nextCai {
			count += 1
			nextCai = cooked[i] + handSpeed
		}
	}

	return count
}
