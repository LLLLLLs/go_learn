// Time        : 2019/07/15
// Description : 平衡二叉树

package avl

import (
	"fmt"
	"math"
)

type AVLNode struct {
	key   int
	value interface{}
	left  *AVLNode
	right *AVLNode
}

func NewNode(k int, v interface{}) *AVLNode {
	return &AVLNode{
		key:   k,
		value: v,
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (a *AVLNode) Get(key int) interface{} {
	if a.key == key {
		return a.value
	}
	if key < a.key {
		if a.left == nil {
			return nil
		}
		return a.left.Get(key)
	} else {
		if a.right == nil {
			return nil
		}
		return a.right.Get(key)
	}
}

func (a *AVLNode) Insert(node *AVLNode) *AVLNode {
	a.insert(node)
	return balance(a)
}

func (a *AVLNode) insert(node *AVLNode) {
	if a.key == node.key {
		a.value = node.value
		return
	}
	if node.key < a.key {
		if a.left == nil {
			a.left = node
			return
		}
		a.left.insert(node)
	} else {
		if a.right == nil {
			a.right = node
			return
		}
		a.right.insert(node)
	}
}

func (a *AVLNode) Remove(key int) *AVLNode {
	return balance(a.remove(key))
}

func (a *AVLNode) remove(key int) *AVLNode {
	if key < a.key && a.left == nil ||
		key > a.key && a.right == nil ||
		key != a.key && a.left == nil && a.right == nil {
		return a
	}
	if a.left == nil && a.right == nil {
		return nil
	}
	if a.key == key {
		if a.right == nil {
			return a.left
		}
		a.right.insert(a.left)
		return a.right
	}
	if key < a.key {
		a.left = a.left.remove(key)
	} else {
		a.right = a.right.remove(key)
	}
	return a
}

// 节点在数的坐标位置(实际上是二位数组的坐标，x标识行，y标识列）
type nodePos struct {
	x      int
	y      int
	origin int // 以根节点对称的左节点y值，用于计算gap
}

// 计算与子节点的间隔
func (p nodePos) gap() int {
	if p.origin%2 == 0 {
		return p.origin / 2
	} else {
		return p.origin/2 + 1
	}
}

// 左孩子坐标
func (p nodePos) left() nodePos {
	return nodePos{
		x:      p.x + p.gap(),
		y:      p.y - p.gap(),
		origin: p.y - p.gap(),
	}
}

// 右孩子坐标
func (p nodePos) right() nodePos {
	return nodePos{
		x:      p.x + p.gap(),
		y:      p.y + p.gap(),
		origin: p.y - p.gap(),
	}
}

func (a *AVLNode) print() {
	depth := depth(a)
	if depth == 1 {
		fmt.Println(a.key)
		return
	}
	// padding = 2 ^ (depth - 2) * 5 + (2 ^ (depth - 2) -1)
	maxPadding := int(math.Pow(float64(2), float64(depth-2)))*5 + int(math.Pow(float64(2), float64(depth-2))-1)
	result := make([][]string, maxPadding/2+1)
	for i := range result {
		result[i] = make([]string, maxPadding)
		for j := range result[i] {
			result[i][j] = " "
		}
	}
	generateTree(a, nodePos{x: 0, y: maxPadding/2 + 1, origin: maxPadding/2 + 1}, result)
	for i := range result {
		for j := range result[i] {
			fmt.Print(result[i][j])
		}
		fmt.Println()
	}
}

func generateTree(node *AVLNode, pos nodePos, result [][]string) {
	result[pos.x][pos.y-1] = fmt.Sprintf("%d", node.key)
	if node.left == nil && node.right == nil {
		return
	}
	if node.left != nil {
		for i := 1; i <= pos.gap()-1; i++ {
			result[pos.x+i][(pos.y - i - 1)] = "/"
		}
		generateTree(node.left, pos.left(), result)
	}
	if node.right != nil {
		for i := 1; i <= pos.gap()-1; i++ {
			result[pos.x+i][(pos.y + i - 1)] = "\\"
		}
		generateTree(node.right, pos.right(), result)
	}
}

func depth(a *AVLNode) int {
	if a == nil {
		return 0
	}
	return max(depth(a.left), depth(a.right)) + 1
}

func llHandle(root *AVLNode) *AVLNode {
	fmt.Println("ll", root.key)
	left := root.left
	root.left = left.right
	left.right = root
	return left
}

func rrHandle(root *AVLNode) *AVLNode {
	fmt.Println("rr", root.key)
	right := root.right
	root.right = right.left
	right.left = root
	return right
}

func lrHandle(root *AVLNode) *AVLNode {
	root.left = rrHandle(root.left)
	return llHandle(root)
}

func rlHandle(root *AVLNode) *AVLNode {
	root.right = llHandle(root.right)
	return rrHandle(root)
}

func balance(root *AVLNode) *AVLNode {
	result, _, _ := balanceDepth(root)
	return result
}

func balanceDepth(root *AVLNode) (result *AVLNode, depth int, left bool) {
	if root == nil {
		return nil, 0, false
	}
	if root.left == nil && root.right == nil {
		return root, 1, false
	}
	lResult, lDept, ll := balanceDepth(root.left)
	root.left = lResult
	rResult, rDept, rl := balanceDepth(root.right)
	root.right = rResult
	if lDept-rDept > 1 {
		if ll {
			result = llHandle(root)
		} else {
			result = lrHandle(root)
		}
		depth = lDept
		left = true
		return
	}
	if rDept-lDept > 1 {
		if rl {
			result = rlHandle(root)
		} else {
			result = rrHandle(root)
		}
		depth = rDept
		left = false
		return
	}
	result = root
	depth = max(lDept, rDept) + 1
	left = lDept > rDept
	return
}
