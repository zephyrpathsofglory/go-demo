package od

/*
题目描述地上共有N个格子，你需要跳完地上所有的格子，但是格子间是有强依赖关系的，跳完前一个格子后，后
续的格子才会被开启，格子间的依赖关系由多组steps数组给出，steps[0]表示前一个格子，steps[1]表示
steps[0]可以开启的格子:

比如[0,1]表示从跳完第0个格子以后第1个格子就开启了，比如[2,1]，[2,3]表示跳完第2个格子后第1个格子和第3
个格子就被开启了。

请你计算是否能由给出的steps数组跳完所有的格子，如果可以输出yes，否则输出no。

说明：

1.你可以从一个格子跳到任意一个开启的格子

2.没有前置依赖条件的格子默认就是开启的

3.如果总数是N，则所有的格子编号为[0,1,2,3…N-1]连续的数组

输入描述输入一个整数N表示总共有多少个格子，接着输入多组二维数组steps表示所有格子之间的依赖关系。

输出描述如果能按照steps给定的依赖顺序跳完所有的格子输出yes，

否则输出no。
*/

// 将 dependencies 看成一个有向图，如果存在 环 就不能跳完。
func canTravelAll(depenencies [][2]int) bool {
	graph := make(map[int][]int, len(depenencies)*2)
	cycleExisted := false

	for _, dep := range depenencies {
		if len(graph[dep[0]]) == 0 {
			graph[dep[0]] = []int{dep[1]}
		} else {
			graph[dep[0]] = append(graph[dep[0]], dep[1])
		}
	}

	for k, v := range graph {
		var f func([]int) bool

		f = func(input []int) bool {
			for _, i := range input {
				if i == k {
					return true
				} else {
					return f(graph[i])
				}
			}

			return false
		}

		cycleExisted = f(v)
		if cycleExisted {
			break
		}
	}

	return !cycleExisted
}
