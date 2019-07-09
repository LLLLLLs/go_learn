// Time        : 2019/07/01
// Description :

package count_say_38

import (
	"strconv"
	"strings"
)

//	The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//	1.     1
//	2.     11		1	=	1	 	=	1个1 			==> 11
//	3.     21		2	=	11		=	2个1 			==> 21
//	4.     1211		3	=	21		=	1个2 1个1		==>	1211
//	5.     111221	4	=	1211	=	1个1 1个2 2个1 	==>	111221
//	1 is read off as "one 1" or 11.
//	11 is read off as "two 1s" or 21.
//	21 is read off as "one 2, then one 1" or 1211.
//
//	Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
//
//	Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//	Example 1:
//
//	Input: 1
//	Output: "1"
//	Example 2:
//
//	Input: 4
//	Output: "1211"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	var count = 0
	var pre = '.'
	var builder = strings.Builder{}
	for _, s := range str {
		if pre == '.' {
			pre = s
			count++
			continue
		}
		if s == pre {
			count++
		} else {
			_, err := builder.WriteString(strconv.Itoa(count))
			OkOrPanic(err)
			_, err = builder.WriteRune(pre)
			OkOrPanic(err)
			pre = s
			count = 1
		}
	}
	_, err := builder.WriteString(strconv.Itoa(count))
	OkOrPanic(err)
	err = builder.WriteByte(str[len(str)-1])
	OkOrPanic(err)
	return builder.String()
}

func OkOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
