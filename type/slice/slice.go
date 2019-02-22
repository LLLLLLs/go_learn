// Time        : 2019/01/18
// Description :

package slice

import "fmt"

func sliceBase() {
	s := make([]int, 5, 10)
	s1 := s[5:10]
	s = append(s, 20, 21, 22)
	s1[0] = 10
	s1[1] = 11
	fmt.Printf("s  = %v,len(s)  = %d,cap(s)  = %d,ptr(s)  = %p\n", s, len(s), cap(s), s)
	fmt.Printf("s1 = %v,len(s1) = %d,cap(s1) = %d,ptr(s1) = %p\n", s1, len(s1), cap(s1), s1)
}

func sliceAppend() {
	sBase := make([]int, 0, 5)
	printInfo(sBase)
	fmt.Println("append 3 integers:1-3")
	sBase = append(sBase, 1, 2, 3)
	printInfo(sBase)
	fmt.Println("append 3 integers:4-6")
	sBase = append(sBase, 4, 5, 6)
	printInfo(sBase)
	fmt.Println("append 4 integers:7-10")
	sBase = append(sBase, 7, 8, 9, 10)
	printInfo(sBase)
	fmt.Println("append 13 integers:11-23")
	sBase = append(sBase, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23)
	printInfo(sBase)
}

func printInfo(s []int) {
	fmt.Printf("%v\tlen:%d\tcap:%d\tptr:%p\n", s, len(s), cap(s), s)
}

func insert(list []int, value, pos int) []int {
	list = append(list[:pos-1], append([]int{value}, list[pos-1:]...)...)
	return list
}
