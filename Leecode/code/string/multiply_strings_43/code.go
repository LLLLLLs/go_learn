// Time        : 2019/07/02
// Description :

package multiply_strings_43

// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
// Example 1:
//
// Input: num1 = "2", num2 = "3"
// Output: "6"
// Example 2:
//
// Input: num1 = "123", num2 = "456"
// Output: "56088"
// Note:
//
// The length of both num1 and num2 is < 110.
// Both num1 and num2 contain only digits 0-9.
// Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// You must not use any built-in BigInteger library or convert the inputs to integer directly.

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var res = make([]byte, 0)
	length := len(num1) + len(num2)
	for i := len(num1) - 1; i >= 0; i-- {
		var carry byte
		for j := len(num2) - 1; j >= 0; j-- {
			index := length - i - j - 2
			tmp := (num1[i]-'0')*(num2[j]-'0') + carry
			if index > len(res)-1 {
				res = append(res, tmp%10+'0')
			} else {
				tmp += res[index] - '0'
				res[index] = tmp%10 + '0'
			}
			carry = tmp / 10
		}
		if carry != 0 {
			res = append(res, carry+'0')
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}
