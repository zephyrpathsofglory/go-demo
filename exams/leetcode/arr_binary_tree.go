package leetcode

import "math"

/*
二叉树也可以用数组来存储，给定一个数组，树的根节点的值储存在下标1，
对于储存在下标n的节点，他的左子节点和右子节点分别储存在下标 2*n 和 2*n+1，
并且我们用-1代表一个节点为空。
给定一个数组存储的二叉树，试求从根节点到最小的叶子节点的路径，路径由节点的值组成。

输入
输入一行为数组的内容，数组的每个元素都是正整数，元素间用空格分割。
注意第一个元素即为根节点的值，即数组的第n元素对应下标 n。
下标0在树的表示中没有使用，所以我们省略了。
输入的树最多为7层。

输出描述
输出从根节点到最小叶子节点的路径上各个节点的值由空格分割
用例保证最小叶子节点只有一个
*/

// nodes[0] = -1， nodes[1] = 根节点
func smallestNodePath(nodes []int) []int {
	frontPointer := 0
	backPointer := 1
	minNode := math.MaxInt
	minNodeIndex := 0
	if len(nodes)%2 == 1 {
		nodes = append(nodes, -1)
	}
	threshold := len(nodes) / 2

	for backPointer <= len(nodes) {
		fpn := nodes[frontPointer]
		bpn := nodes[backPointer]
		if fpn == -1 && bpn == -1 {
			father := nodes[frontPointer/2]

			if father == -1 {
				frontPointer += 2
				backPointer += 2
				continue
			}

			if father < minNode {
				minNode = father
				minNodeIndex = frontPointer / 2
			}
		} else {
			if fpn != -1 && frontPointer >= threshold {
				if fpn < minNode {
					minNode = fpn
					minNodeIndex = frontPointer
				}
			}

			if bpn != -1 && backPointer >= threshold {
				if bpn < minNode {
					minNode = bpn
					minNodeIndex = backPointer
				}
			}
		}

		frontPointer += 2
		backPointer += 2
	}

	path := []int{}
	// 根据 index 推导路径
	for minNodeIndex >= 1 {
		path = append([]int{nodes[minNodeIndex]}, path...)
		minNodeIndex = minNodeIndex / 2
	}

	return path
}
