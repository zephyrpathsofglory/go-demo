package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

func pathSum(root *Node, sum int) (res [][]int) {
	var dfs func(*Node, []int, int)
	dfs = func(root *Node, prefix []int, sum int) {
		if root == nil {
			return
		}

		if isLeaf(root) {
			if sum == root.val {
				res = append(res, append(prefix, root.val))
			}

			return
		}

		dfs(root.left, append(prefix, root.val), sum-root.val)
		dfs(root.right, append(prefix, root.val), sum-root.val)

		return
	}

	dfs(root, []int{}, sum)

	return
}

func isLeaf(n *Node) bool {
	return n.left == nil && n.right == nil
}
