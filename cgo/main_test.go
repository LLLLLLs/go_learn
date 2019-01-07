/*
Author      : lls
Time        : 2018/09/28
Description :
*/

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("init")
	defer fmt.Println("defer")
	m.Run()
}

func TestMain1(t *testing.T) {
	time.Sleep(time.Second)
	fmt.Println("test1")
}

func TestMain2(t *testing.T) {
	fmt.Println("test2")
}

func TestMain3(t *testing.T) {
	fmt.Println("test3")
}
