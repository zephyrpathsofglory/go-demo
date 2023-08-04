package tree

type color bool

const (
	Red   color = true
	Black color = false
)

type RBNode struct {
	Color  color
	Value  int
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
}

type RBTree struct {
	root  *RBNode
	count int
}

func New() *RBTree {
	return &RBTree{
		root: &RBNode{
			Color: Black,
		},
	}
}

func (n *RBNode) grandParent() *RBNode {
	if n.Parent == nil {
		return nil
	}

	return n.Parent.Parent
}

func (n *RBNode) uncleNode() *RBNode {
	if n.Parent == nil {
		return nil
	}

	parent := n.Parent
	if parent.Left == n {
		return parent.Right
	}

	return parent.Left
}

func (t *RBTree) leftRotate(n *RBNode) *RBTree {
	if n.Right == nil {
		return t
	}

	rchild := n.Right

	n.Right = rchild.Left
	if rchild.Left != nil {
		rchild.Left.Parent = n
	}

	rchild.Parent = n.Parent
	rchild.Left = n
	n.Parent = rchild

}
