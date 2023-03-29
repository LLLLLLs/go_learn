package _map

import (
	"fmt"
	"testing"
)

func TestRangeNil(t *testing.T) {
	bigM := map[int]map[int]int{}
	for k, v := range bigM[1] {
		fmt.Println(k, v)
	}
}
