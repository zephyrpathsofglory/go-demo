package tree

type AVLNode struct {
	Height int
	Val    int
	Left   *AVLNode
	Right  *AVLNode
}

func GetHeight(n *AVLNode) int {
	if n == nil {
		return 0
	}

	return n.Height
}

func FindMax(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	if root.Right != nil {
		return FindMax(root.Right)
	}

	return root
}

func FindMin(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		return FindMin(root.Left)
	}

	return root
}

// balance factor, 每个节点都有一个平衡因子，表示该节点的左子树高度减去右子树高度的差值。平衡因子通常在范围[-1, 0, 1]内
// 通过获取节点的平衡因子，可以确定是否需要进行平衡操作来保持树的平衡性。
func GetBF(root *AVLNode) int {
	var lh, rh int

	if root.Left == nil {
		rh = GetHeight(root.Right)
	}
	if root.Right == nil {
		lh = GetHeight(root.Left)
	}

	bf := rh - lh
	if bf < 0 {
		bf = bf * -1
	}

	return bf
}

func LL(root *AVLNode) *AVLNode {
	tmp := root.Left
	root.Left = tmp.Right
	tmp.Right = root

	root.Height = Max(GetHeight(root.Left), GetHeight(root.Right)) + 1
	tmp.Height = Max(GetHeight(tmp.Left), GetHeight(tmp.Right)) + 1

	return tmp
}

func RR(root *AVLNode) *AVLNode {
	tmp := root.Right
	root.Right = tmp.Left
	tmp.Left = root

	root.Height = Max(GetHeight(root.Left), GetHeight(root.Right)) + 1
	tmp.Height = Max(GetHeight(tmp.Left), GetHeight(tmp.Right)) + 1

	return tmp
}

func LR(root *AVLNode) *AVLNode {
	root.Left = RR(root.Left)

	return LL(root)
}

func RL(root *AVLNode) *AVLNode {
	root.Right = LL(root.Right)

	return RR(root)
}

// Max 获取2数中的较大值
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Insert(root *AVLNode, val int) *AVLNode {
	if root == nil {
		root = &AVLNode{
			Val: val,
		}
	} else if val < root.Val {
		root.Left = Insert(root.Left, val)
		if GetHeight(root.Left)-GetHeight(root.Right) == 2 {
			if val < root.Left.Val {
				root = LL(root)
			} else if val > root.Left.Val {
				root = LR(root)
			}
		}
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
		if GetHeight(root.Right)-GetHeight(root.Left) == 2 {
			if val > root.Right.Val {
				root = RR(root)
			} else if val < root.Right.Val {
				root = RL(root)
			}
		}
	}

	root.Height = Max(GetHeight(root.Left), GetHeight(root.Right)) + 1

	return root
}

func Delete(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		if root.Left != nil && root.Right != nil {
			if GetHeight(root.Left) > GetHeight(root.Right) {
				tmp := FindMax(root.Left)
				root.Val = tmp.Val // 将左子树的最大值节点替换根节点
				root.Left = Delete(root.Left, tmp.Val)
			} else {
				tmp := FindMin(root.Right)
				root.Val = tmp.Val
				root.Right = Delete(root.Right, tmp.Val)
			}
		} else {
			if root.Left == nil {
				root = root.Right
			}

			if root.Right == nil {
				root = root.Left
			}
		}

		return root
	} else if root.Val > val {
		root.Left = Delete(root.Left, val)
		// 删除节点后失衡
		if GetHeight(root.Right)-GetHeight(root.Left) == 2 {
			if root.Right.Val > val {
				root = RL(root)
			} else {
				root = RR(root)
			}
		}
	} else if root.Val < val {
		root.Right = Delete(root.Right, val)
		if GetHeight(root.Left)-GetHeight(root.Right) == 2 {
			if root.Left.Val > val {
				root = LL(root)
			} else {
				root = LR(root)
			}
		}
	}

	return root
}
