// Time        : 2019/07/08
// Description :

package primefilter

// 获取第n个素数
func NthPrimeWithList(n int) int {
	primeList := make([]int, 0)
	for i := 0; i < n; i++ {
		addPrime(&primeList)
	}
	return primeList[n-1]
}

func addPrime(list *[]int) {
	var begin = 1
	if len(*list) != 0 {
		begin = (*list)[len(*list)-1]
	}
	for i := begin + 1; ; i++ {
		isPrime := true
		for j := range *list {
			if i%(*list)[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			*list = append(*list, i)
			break
		}
	}
}
