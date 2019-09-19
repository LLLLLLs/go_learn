// Time        : 2019/07/08
// Description : 素数筛 过滤素数

package primefilter

func NthPrimeWithFilter(n int) int {
	prime := make(chan int)
	ctl := make(chan int)
	go PrimeSieve(prime, ctl)
	for i := 0; i < n-1; i++ {
		_ = <-prime
	}
	result := <-prime
	close(ctl)
	return result
}

func nextNum(out, close chan int) {
	for i := 2; ; i++ {
		select {
		case <-close:
			return
		default:
			out <- i
		}
	}
}

func PrimeFilter(prime int, in, out, close chan int) {
	for {
		select {
		case <-close:
			return
		case num := <-in:
			if num%prime != 0 {
				out <- num
			}
		}
	}
}

func PrimeSieve(out, close chan int) {
	c := make(chan int)
	go nextNum(c, close)
	for {
		select {
		case <-close:
			return
		case prime := <-c:
			out <- prime
			newC := make(chan int)
			go PrimeFilter(prime, c, newC, close)
			c = newC
		}
	}
}
