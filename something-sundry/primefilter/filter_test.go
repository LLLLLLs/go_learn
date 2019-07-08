// Time        : 2019/07/08
// Description :

package primefilter

import (
	"fmt"
	"testing"
)

func TestPrimeFilter(t *testing.T) {
	prime := make(chan int)
	ctl := make(chan int)
	go PrimeSieve(prime, ctl)
	for i := 0; i < 10; i++ {
		fmt.Println(<-prime)
	}
	close(ctl)
	fmt.Println(<-prime)
}

func TestNthPrimeWithFilter(t *testing.T) {
	fmt.Println(NthPrimeWithFilter(1))
	fmt.Println(NthPrimeWithFilter(2))
	fmt.Println(NthPrimeWithFilter(3))
	fmt.Println(NthPrimeWithFilter(4))
	fmt.Println(NthPrimeWithFilter(5))
	fmt.Println(NthPrimeWithFilter(6))
	fmt.Println(NthPrimeWithFilter(7))
	fmt.Println(NthPrimeWithFilter(8))
}
