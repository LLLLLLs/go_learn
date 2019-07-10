// Time        : 2019/07/09
// Description :

package addbinary_67

// Given two binary strings, return their sum (also a binary string).
//
// The input strings are both non-empty and contains only characters 1 or 0.
//
// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"

func addBinary(a string, b string) string {
	var carry = byte(0)
	var result = make([]byte, 0)
	for i := 0; ; i++ {
		if i >= len(a) && i >= len(b) {
			break
		}
		m, n := byte(0), byte(0)
		if i < len(a) {
			m = a[len(a)-i-1] - '0'
		}
		if i < len(b) {
			n = b[len(b)-i-1] - '0'
		}
		mid := m + n + carry
		carry = mid / 2
		result = append([]byte{mid%2 + '0'}, result...)
	}
	if carry != 0 {
		result = append([]byte{carry + '0'}, result...)
	}
	return string(result)
}
