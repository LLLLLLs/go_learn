// Time        : 2019/07/18
// Description :

package unique_binary_search_II_95

import . "go_learn/leecode/tree/base"

//Input: 3
//Output:
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//Explanation:
//The above output corresponds to the 5 unique BST's shown below:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	nodes := make([]int, n)
	for i := range nodes {
		nodes[i] = i + 1
	}
	return recursive(nodes)
}

type key struct {
	begin, end int
}

var cache = make(map[key][]*TreeNode)

func recursive(nodes []int) (result []*TreeNode) {
	if len(nodes) == 0 {
		return []*TreeNode{nil}
	}
	if c, ok := cache[key{nodes[0], nodes[len(nodes)-1]}]; ok {
		return c
	}
	if len(nodes) == 1 {
		return []*TreeNode{{Val: nodes[0]}}
	}
	for i := range nodes {
		left := recursive(nodes[:i])
		right := recursive(nodes[i+1:])
		mid := make([]*TreeNode, 0, len(left)*len(right))
		for m := range left {
			for n := range right {
				mid = append(mid, &TreeNode{Val: nodes[i], Left: left[m], Right: right[n]})
			}
		}
		result = append(result, mid...)
	}
	cache[key{nodes[0], nodes[len(nodes)-1]}] = result
	return
}
