package randomutils

func LCGDice(randNum int, arr ...int) func() int {
	//随机数表达式，闭包实现
	var param = [3]int{8121, 28411, 134456}
	for i := range arr {
		param[i] = arr[i]
	}
	return func() int {
		randNum = (8121*randNum + 28411) % 134456
		return randNum
	}
}

func InitDice() (int, func() int) {
	//随机得到一个种子，返回给客户端，并生成随机数生成器
	seed := RandomInt(0, 65536)
	dice := LCGDice(seed)
	return seed, dice
}

func InitDiceWithSeed(seed int) func() int {
	dice := LCGDice(seed)
	return dice
}
