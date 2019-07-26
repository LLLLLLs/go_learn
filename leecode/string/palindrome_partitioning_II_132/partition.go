// Time        : 2019/07/25
// Description :

package palindrome_partitioning_II_132

//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return the minimum cuts needed for a palindrome partitioning of s.
//
//Example:
//
//Input: "aab"
//Output: 1
//Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

func _minCut(s string) int {
	m := 0
	palindromeMap := make(map[int]map[int]bool) // weather i to j is a palindrome
	minChar := make(map[int]int)                // mini cut for 0 to i string

	for i := 0; i < len(s); i++ {
		m = i
		for j := 0; j <= i; j++ {
			if s[j:j+1] == s[i:i+1] {
				_, v := palindromeMap[j+1][i-1]
				if v == true || j+1 > i-1 {
					palindromeMapChild := make(map[int]bool)
					palindromeMapChild[i] = true
					palindromeMap[j] = palindromeMapChild
					if j == 0 {
						m = 0
					} else {
						if m > (minChar[j-1] + 1) {
							m = minChar[j-1] + 1
						}
					}
				}
			}
		}
		minChar[i] = m
	}
	return minChar[len(s)-1]
}

func minCut(s string) int {
	maxPalindrome := make(map[int]map[int]bool)
	minC := make(map[int]int)
	for i := 0; i < len(s); i++ {
		cut := i
		for j := 0; j <= i; j++ {
			if s[j] == s[i] {
				_, v := maxPalindrome[j+1][i-1]
				if v || j+1 > i-1 {
					maxPalindrome[j] = map[int]bool{i: true}
					if j == 0 {
						cut = 0
					} else if cut > minC[j-1]+1 {
						cut = minC[j-1] + 1
					}
				}
			}
		}
		minC[i] = cut
	}
	return minC[len(s)-1]
}

func minCut2(s string) int {
	mc := make(map[int]int)
	for i := range s {
		mc[i] = i
	}
	mc[-1] = -1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(s); i++ {
		for j := 0; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			mc[i+j] = min(mc[i+j], mc[i-j-1]+1)
		}
		for j := 1; i-j+1 >= 0 && i+j < len(s) && s[i-j+1] == s[i+j]; j++ {
			mc[i+j] = min(mc[i+j], mc[i-j]+1)
		}
	}
	return mc[len(s)-1]
}

func minCut3(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	pal := make([][]bool, length)
	for i := range pal {
		pal[i] = make([]bool, length)
	}
	cut := make([]int, length+1)
	cut[length] = -1
	for i := length - 1; i >= 0; i-- {
		cut[i] = length - i
		for j := i; j < length; j++ {
			if i == j || (s[i] == s[j] && (i+1 >= j-1 || pal[i+1][j-1])) {
				pal[i][j] = true
				if cut[i] > cut[j+1]+1 {
					cut[i] = cut[j+1] + 1
				}
			}
		}
	}
	return cut[0]
}
