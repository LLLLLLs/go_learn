// Time        : 2019/07/17
// Description :

package pickrocket

// 捡石头
// 地上有n块石头，每块石头都有积分
// 双方轮流捡石块(我方先捡)，下一个人最多可以捡上一个人捡的2倍，最少捡1块
// 求我方最高积分

func pickRock(_ int, scores []int) int {
	a, _ := backtrack(scores, 1, 1)
	return a
}

type key struct {
	length, last int
}

type value struct {
	a, b int
}

var result = make(map[key]value)

func backtrack(scores []int, index, last int) (a, b int) {
	if len(scores) < last*2 {
		total := 0
		for i := range scores {
			total += scores[i]
		}
		if index%2 == 1 {
			return total, 0
		} else {
			return 0, total
		}
	}
	if v, ok := result[key{length: len(scores), last: last}]; ok {
		if index%2 == 1 {
			return v.a, v.b
		} else {
			return v.b, v.a
		}
	}
	this := 0
	max := 0
	for pick := 0; pick < last*2; pick++ {
		this += scores[pick]
		na, nb := backtrack(scores[pick+1:], index+1, pick+1)
		if index%2 == 1 {
			if na+this > max {
				max = na + this
				a = max
				b = nb
			}
		} else {
			if nb+this > max {
				max = nb + this
				b = max
				a = na
			}
		}
	}
	v := value{a, b}
	if index%2 == 0 {
		v = value{b, a}
	}
	result[key{len(scores), last}] = v
	return
}
