package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
