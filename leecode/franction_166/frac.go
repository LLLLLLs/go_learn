//@author: lls
//@time: 2020/08/27
//@desc:

package franction_166

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	integer := ""
	if numerator*denominator < 0 {
		integer = "-"
	}
	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}
	i := numerator / denominator
	integer += strconv.Itoa(i)
	numerator %= denominator
	if numerator == 0 {
		return integer
	} else {
		integer += "."
	}
	sub := make(map[int]int)
	fraction := make([]byte, 0)
	for numerator != 0 {
		if l, repeated := sub[numerator]; repeated {
			rpd := make([]byte, len(fraction)-l)
			copy(rpd, fraction[l:])
			fraction = append(fraction[:l], '(')
			fraction = append(fraction, rpd...)
			fraction = append(fraction, ')')
			break
		}
		sub[numerator] = len(fraction)
		numerator *= 10
		r := numerator / denominator
		fraction = append(fraction, byte(r)+'0')
		numerator %= denominator
	}
	return integer + string(fraction)
}
