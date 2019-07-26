// Time        : 2019/07/23
// Description :

package word_ladder_II_126

// Given two words (beginWord and endWord), and a dictionary's word list,
// find all shortest transformation sequence(s) from beginWord to endWord, such that:
//
// Only one letter can be changed at a time
// Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
// Note:
//
// Return an empty list if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// You may assume no duplicates in the word list.
// You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:
//
// Input:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
//
// Output:
// [
//   ["hit","hot","dot","dog","cog"],
//   ["hit","hot","lot","log","cog"]
// ]
// Example 2:
//
// Input:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
//
// Output: []
//
// Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func myFindLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return nil
	}
	result := make([][]string, 0)
	queue := [][]string{{beginWord}}
	visited := make(map[string]struct{})
	level := 1
	length := len(wordList) + 1
	for len(queue) != 0 {
		ladder := queue[0]
		queue = queue[1:]
		if len(ladder) > level {
			for w := range visited {
				delete(words, w)
			}
			visited = make(map[string]struct{})
			level = len(ladder)
		}
		if len(ladder) > length {
			break
		} else {
			if ladder[len(ladder)-1] == endWord {
				result = append(result, ladder)
				length = len(ladder)
				continue
			}
		}
		prev := ladder[len(ladder)-1]
		for i := range prev {
			for j := 'a'; j <= 'z'; j++ {
				newWord := prev[:i] + string(j) + prev[i+1:]
				if _, ok := words[newWord]; !ok || newWord == prev {
					continue
				}
				visited[newWord] = struct{}{}
				queue = append(queue, append(append([]string{}, ladder...), newWord))
			}
		}
	}
	return result
}

// leetcode最优解
func findLadders2(beginWord string, endWord string, wordList []string) [][]string {
	wordDict := make(map[string]struct{})
	for _, word := range wordList {
		wordDict[word] = struct{}{}
	}
	// 提前返回
	if _, ok := wordDict[endWord]; !ok {
		return [][]string{}
	}

	src, dst := map[string]struct{}{}, map[string]struct{}{}
	src[beginWord] = struct{}{}
	dst[endWord] = struct{}{}

	found := false
	paths := make(map[string][]string)
	backward := false

	for len(src) != 0 && len(dst) != 0 && !found {

		// 始终遍历 小数组
		if len(src) > len(dst) {
			src, dst = dst, src
			backward = !backward
		}
		for w := range src {
			delete(wordDict, w)
		}

		newSrc := make(map[string]struct{})
		for word := range src {
			bytes := []byte(word) // 转成 []byte，然后修改值
			for i := 0; i < len(bytes); i++ {
				for j := 0; j < 26; j++ {
					bytes[i] = byte(j) + 'a' // 修改为 a-z
					source := word
					target := string(bytes)       // 转成 string
					if _, ok := dst[target]; ok { // 已经连通
						if backward {
							source, target = target, source
						}
						if paths[target] == nil {
							paths[target] = make([]string, 0)
						}
						paths[target] = append(paths[target], source)
						found = true
					} else {
						if _, ok := wordDict[target]; ok {
							newSrc[target] = struct{}{}
							if backward {
								source, target = target, source
							}
							if paths[target] == nil {
								paths[target] = make([]string, 0)
							}
							paths[target] = append(paths[target], source)
						}
					}
				}
				// 恢复原来的值
				bytes[i] = word[i]
			}
		}
		src = newSrc
	}
	ans := make([][]string, 0)
	getPath(paths, beginWord, endWord, []string{endWord}, &ans)

	return ans
}
func getPath(parents map[string][]string, beginWord, cur string, path []string, ans *[][]string) {
	if cur == beginWord {
		newPath := make([]string, len(path))
		copy(newPath, path)
		for i, j := 0, len(newPath)-1; i < j; i, j = i+1, j-1 {
			newPath[i], newPath[j] = newPath[j], newPath[i]
		}
		*ans = append(*ans, newPath)
		return
	}

	for _, p := range parents[cur] {
		path = append(path, p)
		getPath(parents, beginWord, p, path, ans)
		path = path[:len(path)-1]
	}
}

// 仿上述解实现
// 				lot → log
// 			 ↗			  ↘
// hit → hot 				 cog
// 			 ↘			  ↗
// 				dot → dog
//
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return nil
	}
	begin := map[string]struct{}{beginWord: {}}
	end := map[string]struct{}{endWord: {}}
	path := make(map[string][]string)
	linked := false
	back := false

	for len(begin) != 0 && len(end) != 0 && !linked {
		if len(begin) > len(end) {
			begin, end = end, begin
			back = !back
		}
		newBegin := make(map[string]struct{})
		for w := range begin {
			delete(words, w)
		}
		for w := range begin {
			for j := range w {
				for i := 'a'; i <= 'z'; i++ {
					source := w
					newWord := w[:j] + string(i) + w[j+1:]
					if _, ok := words[newWord]; ok {
						if _, ok2 := end[newWord]; ok2 {
							// 连通
							linked = true
						}
						newBegin[newWord] = struct{}{}
						if back {
							source, newWord = newWord, source
						}
						if path[newWord] == nil {
							path[newWord] = make([]string, 0)
						}
						path[newWord] = append(path[newWord], source)
					}
				}
			}
		}
		begin = newBegin
	}
	result := make([][]string, 0)
	genPaths(path, beginWord, []string{endWord}, &result)
	return result
}

func genPaths(paths map[string][]string, begin string, ladder []string, result *[][]string) {
	if ladder[len(ladder)-1] == begin {
		ladderCopy := make([]string, len(ladder))
		copy(ladderCopy, ladder)
		for i := 0; i < len(ladderCopy)/2; i++ {
			ladderCopy[i], ladderCopy[len(ladderCopy)-i-1] = ladderCopy[len(ladderCopy)-i-1], ladderCopy[i]
		}
		*result = append(*result, ladderCopy)
		return
	}
	for _, w := range paths[ladder[len(ladder)-1]] {
		ladder = append(ladder, w)
		genPaths(paths, begin, ladder, result)
		ladder = ladder[:len(ladder)-1]
	}
}
