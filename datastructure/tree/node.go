package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)

		traversal(root.Left)
		traversal(root.Right)
	}

	traversal(root)

	return
}

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(*TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		traversal(root.Right)

		res = append(res, root.Val)
	}

	traversal(root)

	return
}

// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(*TreeNode)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}

	traversal(root)

	return
}

// 层序遍历, 广度优先遍历
func levelTraversal(root *TreeNode) (res [][]int) {
	var traversal func(*TreeNode, int)
	traversal = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], root.Val)

		traversal(root.Left, depth+1)
		traversal(root.Right, depth+1)
	}

	traversal(root, 0)

	return
}

// 二叉树的右视图
func rightSideView(root *TreeNode) (res []int) {
	var traversal func(*TreeNode, int)
	traversal = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(res) == depth {
			res = append(res, root.Val)
		} else {
			res[depth] = root.Val
		}

		traversal(root.Left, depth+1)
		traversal(root.Right, depth+1)
	}

	traversal(root, 0)

	return
}

// 翻转二叉树
func invert(root *TreeNode) {
	var invert func(*TreeNode)

	invert = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left

		invert(root.Right)
		invert(root.Left)
	}

	invert(root)

	return
}
