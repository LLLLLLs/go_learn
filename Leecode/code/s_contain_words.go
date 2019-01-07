/*
Author      : lls
Time        : 2018/10/30
Description :
*/

package code

//给定一个字符串 s 和一些长度相同的单词 words。在 s 中找出可以恰好串联 words 中所有单词的子串的起始位置。
//
//注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//
//示例 1:
//
//输入:
//  s = "barfoothefoobarman",
//  words = ["foo","bar"]
//输出: [0,9]
//解释: 从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。

type wordTree struct {
	father   *wordTree
	children map[rune]*wordTree
	node     rune
	reached  bool
}

func newWordTree(node rune) *wordTree {
	return &wordTree{children: make(map[rune]*wordTree), node: node}
}

func WordsContainerWithTree(s string, words []string) []int {
	wordLen := len(words[0])
	wordCount := len(words)
	head := newWordTree(0)
	for _, word := range words {
		prev := head
		for i, r := range []rune(word) {
			child, ok := prev.children[r]
			if !ok {
				child = newWordTree(r)
				child.father = prev
				if i == len([]rune(word))-1 {
					child.children = nil
				}
				prev.children[r] = child
			}
			prev = child
		}
	}
	runes := []rune(s)
	result := make([]int, 0)
	for i := range runes {
		if len(runes)-i < wordLen*wordCount {
			break
		}
		if match(runes[i:i+wordLen*wordCount], head) == wordCount {
			result = append(result, i)
		}
	}
	return result
}

func match(s []rune, tree *wordTree) int {
	modified := make([]*wordTree, 0)
	count := 0
	prev := tree
	for _, r := range s {
		if child, ok := prev.children[r]; ok {
			// 到达叶子节点
			if child.children == nil {
				// 该叶子未到达过 计数+1 接下来继续从根遍历
				if child.reached == false {
					count++
					prev = tree
					child.reached = true
					modified = append(modified, child)
					continue
				} else {
					// 该叶子已到达，重复遍历，失败
					break
				}
			} else {
				// 未到达叶子，继续遍历
				prev = child
			}
		} else {
			// 无该节点，遍历失败
			break
		}
	}
	// 重置树状态
	for _, t := range modified {
		t.reached = false
	}
	return count
}
