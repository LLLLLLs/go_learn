// Time        : 2019/01/22
// Description :

package merge_k_linked_list_23

import (
	"go_learn/Leecode/code/linked_list/base"
	"math/rand"
	"testing"
	"time"
)

var lists = func() []*base.ListNode {
	l1 := base.NewList(1)
	l1.Add(4, 5)
	l2 := base.NewList(1)
	l2.Add(3, 4)
	l3 := base.NewList(2)
	l3.Add(6)
	return []*base.ListNode{l1, l2, l3}
}()

var listBench = func() (list []*base.ListNode) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 90; i++ {
		l := base.NewList(1)
		num := 1
		for i := 0; i < 90; i++ {
			num += rand.Intn(2)
			l.Add(num)
		}
		list = append(list, l)
	}
	//for i := 0; i < 9; i++ {
	//	list[i].Print()
	//}
	return
}()

func TestMergeKLists1(t *testing.T) {
	l := mergeKLists1(lists)
	l.Print()
}

func TestMergeKLists2_1(t *testing.T) {
	l := mergeKLists2_1(lists)
	l.Print()
}

func TestMergeKLists2_2(t *testing.T) {
	l := mergeKLists2_2(listBench)
	l.Print()
}

func TestMergeKLists3(t *testing.T) {
	l := mergeKLists3(lists)
	l.Print()
}

func TestMergeKLists5(t *testing.T) {
	l := mergeKLists5(lists)
	l.Print()
}

func BenchmarkList1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists1(listBench)
	}
}

func BenchmarkList2_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists2_1(listBench)
	}
}

func BenchmarkList2_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists2_2(listBench)
	}
}

func BenchmarkList3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists3(listBench)
	}
}

func BenchmarkList4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeKLists4(listBench)
	}
}
