package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := preorderTraversal(tree)
	assert.Equal(t, []int{10, 6, 3, 8, 15, 12, 18}, res)
}

func TestPostorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := postorderTraversal(tree)
	assert.Equal(t, []int{3, 8, 6, 12, 18, 15, 10}, res)
}

func TestInorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := inorderTraversal(tree)
	assert.Equal(t, []int{3, 6, 8, 10, 12, 15, 18}, res)
}

func TestLevelTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := levelTraversal(tree)
	assert.Equal(t, [][]int{
		{10},
		{6, 15},
		{3, 8, 12, 18},
	}, res)
}

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	res := rightSideView(tree)
	assert.Equal(t, []int{10, 15, 18}, res)
}

func TestInvert(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	invert(tree)
	assert.Equal(t, 10, tree.Val)
	assert.Equal(t, 15, tree.Left.Val)
	assert.Equal(t, 18, tree.Left.Left.Val)
	assert.Equal(t, 12, tree.Left.Right.Val)
	assert.Equal(t, 6, tree.Right.Val)
	assert.Equal(t, 8, tree.Right.Left.Val)
	assert.Equal(t, 3, tree.Right.Right.Val)
}
