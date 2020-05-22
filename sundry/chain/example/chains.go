//@author: lls
//@time: 2020/05/19
//@desc:

package example

import (
	"fmt"
	"golearn/sundry/chain"
	"golearn/util/randutil"
)

type Key string

const TestKey = "test"

type Generate struct{}

func (g Generate) Handle(c chain.Context) error {
	value := randutil.RandInt(1, 10)
	fmt.Println("随机数=", value)
	c.SetValue(TestKey, value)
	return c.Next()
}

type Add struct{}

func (a Add) Handle(c chain.Context) error {
	value := c.Value(TestKey).(int)
	fmt.Println("Add得到的数=", value)
	value += 20
	c.SetValue(TestKey, value)
	fmt.Println("第一次Add:x + 20 =", value)
	err := c.Next()
	if err != nil {
		return err
	}
	value = c.Value(TestKey).(int)
	fmt.Println("调用链回到Add时=", value)
	value += 30
	c.SetValue(TestKey, value)
	fmt.Println("第二次Add:x + 30 =", value)
	return nil
}

type Multi struct{}

func (m Multi) Handle(c chain.Context) error {
	value := c.Value(TestKey).(int)
	fmt.Println("Multi得到的数=", value)
	value *= 2
	c.SetValue(TestKey, value)
	fmt.Println("Multi结果:x * 2 =", value)
	return c.Next()
}
