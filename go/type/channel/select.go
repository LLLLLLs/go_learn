//@author: lls
//@time: 2020/05/20
//@desc:

package channel

func SelectDeadlock() {
	a, b := make(chan int), make(chan int)
	a <- 1
	for {
		select {
		case <-a:
			b <- 1
		case <-b:
			a <- 1
		}
	}
}
