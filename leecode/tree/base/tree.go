// Time        : 2019/07/12
// Description :

package base

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// 中序遍历 左中右
func (t *TreeNode) InorderTraversal() []int {
	io := make([]int, 0)
	inorder(t, &io)
	return io
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}

// 先序遍历 中左右
func (t *TreeNode) PreorderTraversal() []int {
	io := make([]int, 0)
	preorder(t, &io)
	return io
}

func preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
}

// 后序遍历 左右中
func (t *TreeNode) PostorderTraversal() []int {
	io := make([]int, 0)
	postorder(t, &io)
	return io
}

func postorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, result)
	postorder(root.Right, result)
	*result = append(*result, root.Val)
}
