// Time        : 2019/07/15
// Description : 平衡二叉树 //TODO

package avl

import "fmt"

type AVLNode struct {
	key    int
	value  interface{}
	height int
	left   *AVLNode
	right  *AVLNode
}

func NewNode(k int, v interface{}) *AVLNode {
	return &AVLNode{
		key:    k,
		value:  v,
		height: 1,
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (a *AVLNode) Get(key int) interface{} {
	if a == nil {
		return nil
	}
	if a.key == key {
		return a.value
	}
	if key < a.key {
		return a.left.Get(key)
	} else {
		return a.right.Get(key)
	}
}

func (a *AVLNode) PrintInorder() {
	inorder(a)
	fmt.Printf("nil\n")
}

func inorder(root *AVLNode) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%v-->", root.value)
	inorder(root.right)
}
