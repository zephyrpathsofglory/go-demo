// https://leetcode.cn/problems/shortest-path-visiting-all-nodes/solutions/918311/fang-wen-suo-you-jie-dian-de-zui-duan-lu-mqc2/
package bfs

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
)

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}

	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			fmt.Println(size.Of(seen))
			fmt.Println(size.Of(q))
			return t.dist
		}
		// 搜索相邻的节点
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}

// 用 map 来表示 seen，更加的耗费内存，seen[i] 的 len 没有达到 2 ** n, 但是也不小，使用map更加耗费内存，同时大量存在
// u&mask 相同，但是 dist 不同的情况，所以不应该在 dist 发生变化后将 seen 重置，导致大量 mask 信息丢失后，q 的长度大增。
//
// func shortestPathLength2(graph [][]int) int {
// 	n := uint8(len(graph))
// 	type tuple struct {
// 		u    uint8
// 		mask int
// 		dist uint8
// 	}
// 	q := []tuple{}
// 	seen := make([]map[int]bool, n)
// 	for i := range seen {
// 		seen[i] = make(map[int]bool)
// 		q = append(q, tuple{uint8(i), 1 << i, 0})
// 	}

// 	// currentDist := 0

// 	for {
// 		t := q[0]
// 		q = q[1:]
// 		if t.mask == 1<<n-1 {
// 			fmt.Println(size.Of(seen))
// 			fmt.Println(size.Of(q))

// 			return int(t.dist)
// 		}

// 		// d := t.dist
// 		// if d > uint8(currentDist) {
// 		// 	for i := range seen {
// 		// 		seen[i] = make(map[int]bool, n)
// 		// 	}
// 		// 	currentDist = int(d)
// 		// }

// 		// 搜索相邻的节点
// 		for _, v := range graph[t.u] {
// 			maskV := t.mask | 1<<v
// 			if !seen[v][maskV] {
// 				q = append(q, tuple{uint8(v), maskV, t.dist + 1})
// 				seen[v][maskV] = true
// 			}
// 		}
// 	}
// }
