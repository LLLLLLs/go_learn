// Time        : 2019/09/24
// Description :

package randutil

import (
	"fmt"
	"testing"
)

func TestDice(t *testing.T) {
	dice1 := InitDiceWithSeed(12345)
	for i := 0; i < 10; i++ {
		fmt.Println(dice1.Roll(0, 7))
	}
	diceRand := InitDice()
	var result = make(map[int]int)
	for i := 0; i < 1000000; i++ {
		result[diceRand.Roll(0, 7)]++
	}
	fmt.Printf("%+v\n", result)

}
