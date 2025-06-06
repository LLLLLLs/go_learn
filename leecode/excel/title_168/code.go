package title_168

// 给你一个整数columnNumber ，返回它在 Excel 表中相对应的列名称。
//
// 例如：
//
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
//
// 示例 1：
//
// 输入：columnNumber = 1
// 输出："A"
// 示例 2：
//
// 输入：columnNumber = 28
// 输出："AB"
// 示例 3：
//
// 输入：columnNumber = 701
// 输出："ZY"
// 示例 4：
//
// 输入：columnNumber = 2147483647
// 输出："FXSHRXW"
//
// 提示：
//
// 1 <= columnNumber <= 231 - 1

func convertToTitle(columnNumber int) string {
	cpy := columnNumber
	length := int8(0)
	for cpy != 0 {
		length++
		cpy = (cpy - 1) / 26
	}
	chars := make([]byte, length)
	for columnNumber != 0 {
		length--
		chars[length] = 'A' + byte((columnNumber-1)%26)
		columnNumber = (columnNumber - 1) / 26
	}
	return string(chars)
}
