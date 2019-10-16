// Time        : 2019/09/24
// Description : 自定义随机函数

package myrand

import "golearn/util"

type Dice interface {
	// 获取闭区间[min,max]中的随机数
	Roll(min, max int) int
	// 获取随机种子
	Seed() int
}

type dice struct {
	f    func(n int) int
	seed int
}

func (d *dice) Roll(min, max int) int {
	if min > max {
		panic("max must greater than min")
	}
	rand := d.f(max - min + 1)
	d.seed = rand
	return rand%(max-min+1) + min
}

func (d *dice) Seed() int {
	return d.seed
}

func InitDice() Dice {
	//随机得到一个种子，返回给客户端，并生成随机数生成器
	seed := util.RandInt(0, 65536)
	return InitDiceWithSeed(seed)
}

func InitDiceWithSeed(seed int) Dice {
	f := LCGDice(seed)
	return &dice{
		f:    f,
		seed: seed,
	}
}

func LCGDice(randNum int, arr ...int) func(n int) int {
	//随机数表达式，闭包实现
	var param = [3]int{8121, 28411, 134456}
	for i := range arr {
		if i > 2 {
			break
		}
		param[i] = arr[i]
	}
	return func(n int) int {
		randNum = (param[0]*randNum + param[1]) % param[2]
		max := param[2] - param[2]%n - 1
		for randNum > max {
			randNum = (param[0]*randNum + param[1]) % param[2]
		}
		return randNum
	}
}
