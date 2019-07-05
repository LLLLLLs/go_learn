// Time        : 2019/07/04
// Description :

package valid_num_65

import "strings"

// Validate if a given string can be interpreted as a decimal number.
//
// Some examples:
// "0" => true
// " 0.1 " => true
// "abc" => false
// "1 a" => false
// "2e10" => true
// " -90e3   " => true
// " 1e" => false
// "e3" => false
// " 6e-1" => true
// " 99e2.5 " => false
// "53.5e93" => true
// " --6 " => false
// "-+3" => false
// "95a54e53" => false
//
// Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one. However, here is a list of characters that can be in a valid decimal number:
//
// Numbers 0-9
// Exponent - "e"
// Positive/negative sign - "+"/"-"
// Decimal point - "."
// Of course, the context of these characters also matters in the input.

func isNumber(s string) bool {
	s = strings.Trim(s, string(" "))
	if len(s) == 0 {
		return false
	}
	prefixOk, eExit, dot := false, false, false
	for i, ch := range s {
		switch ch {
		case 'e':
			if eExit || !prefixOk {
				return false
			}
			eExit, prefixOk = true, false
		case '.':
			if eExit || dot {
				return false
			}
			if !((i < len(s)-1 && (s[i+1] >= '0' && s[i+1] <= '9')) ||
				(i > 0 && (s[i-1] >= '0' && s[i-1] <= '9'))) {
				return false
			}
			dot = true
		case '-', '+':
			if i != 0 && s[i-1] != 'e' {
				return false
			}
		default:
			if ch < '0' || ch > '9' {
				return false
			} else {
				prefixOk = true
			}
		}
	}
	return prefixOk
}
