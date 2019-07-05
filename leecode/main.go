/*
Author      : lls
Time        : 2018/10/29
Description :
*/

package main

import (
	"fmt"
	"go_learn/leecode/code"
)

func main() {
	s := sudoku{}
	s.solve()
}

type lru struct {
	name string
}

type sudoku struct {
}

func (lru lru) cache() {
	cache := code.NewLRUCache(5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	fmt.Println(cache.Get(1))
	cache.Put(2, 6)
	cache.Put(3, 7)
	cache.Put(4, 8)
	cache.Put(5, 9)
	fmt.Println(cache.Get(2))
	cache.Put(6, 10)
	fmt.Println(cache.Get(1))
}

var Sudoku = [9][9]int{
	{0, 0, 0, 0, 0, 0, 8, 0, 0},
	{0, 8, 2, 4, 0, 0, 0, 0, 0},
	{1, 9, 0, 0, 6, 3, 0, 0, 0},
	{0, 5, 0, 0, 8, 0, 7, 0, 0},
	{6, 7, 8, 2, 0, 9, 1, 4, 3},
	{0, 0, 3, 0, 4, 0, 0, 8, 0},
	{0, 0, 0, 6, 2, 0, 0, 9, 4},
	{0, 0, 0, 0, 0, 5, 6, 1, 0},
	{0, 0, 0, 6, 0, 0, 0, 0, 0}}

func (sud sudoku) solve() {
	code.SolveSudoku(Sudoku)
}
