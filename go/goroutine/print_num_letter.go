// Time        : 2019/08/23
// Description :

package goroutine

import "fmt"

func printNumAndLetter() {
	letter, number := make(chan struct{}), make(chan struct{})
	letterDone, numberDone := make(chan struct{}), make(chan struct{})
	go func() {
		for i := 1; i < 28; i += 2 {
			select {
			case <-number:
				fmt.Print(i)
				fmt.Print(i + 1)
				letter <- struct{}{}
			}
		}
		numberDone <- struct{}{}
	}()
	go func() {
		for x := 'A'; x <= 'Z'; x += 2 {
			select {
			case <-letter:
				fmt.Print(string(x))
				fmt.Print(string(x + 1))
				number <- struct{}{}
			}
		}
		<-letter
		letterDone <- struct{}{}
	}()
	number <- struct{}{}
	<-letterDone
	<-numberDone
}
