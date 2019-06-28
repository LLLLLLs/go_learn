// Time        : 2019/06/26
// Description :

package strStr_28

// Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Example 1:
//
// Input: haystack = "hello", needle = "ll"
// Output: 2
//
// Example 2:
//
// Input: haystack = "aaaaa", needle = "bba"
// Output: -1

func strStr(haystack string, needle string) int {
	//return strings.Index(haystack, needle)
	if haystack == needle || needle == "" {
		return 0
	}
	for i := range haystack {
		if haystack[i] == needle[0] {
			if i+len(needle) > len(haystack) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
