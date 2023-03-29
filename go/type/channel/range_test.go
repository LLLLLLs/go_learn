package channel

import "testing"

func TestRange(t *testing.T) {
	c := make(chan int, 100)
	go Range(c, 1)
	go Range(c, 2)
	go ForSelect(c, 3)
	Send(c)
}
