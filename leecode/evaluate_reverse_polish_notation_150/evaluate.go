// Time        : 2019/07/30
// Description :

package evaluate_reverse_polish_notation_150

import "strconv"

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
//
// Valid operators are +, -, *, /. Each operand may be an integer or another expression.
//
// Note:
//
// Division between two integers should truncate toward zero.
// The given RPN expression is always valid.
// That means the expression would always evaluate to a result and there won't be any divide by zero operation.
// Example 1:
//
// Input: ["2", "1", "+", "3", "*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9
// Example 2:
//
// Input: ["4", "13", "5", "/", "+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6
// Example 3:
//
// Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// Output: 22
// Explanation:
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	pop := func() int {
		num := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return num
	}
	push := func(num int) {
		stack = append(stack, num)
	}
	for i := range tokens {
		switch tokens[i] {
		case "+":
			push(pop() + pop())
		case "-":
			num1, num2 := pop(), pop()
			push(num2 - num1)
		case "*":
			push(pop() * pop())
		case "/":
			num1, num2 := pop(), pop()
			push(num2 / num1)
		default:
			num, _ := strconv.Atoi(tokens[i])
			push(num)
		}
	}
	return pop()
}
