// Time        : 2019/06/25
// Description :

package _map

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	var m = make(map[int16]bool)
	m[1] = true
	m[2] = true
	fmt.Println(len(m))
	delete(m, 1)
	fmt.Println(len(m))
	delete(m, 2)
	fmt.Println(len(m))
}

func TestCopyMap(t *testing.T) {
	m1 := map[string]bool{
		"1": true,
		"2": false,
	}
	m2 := m1
	delete(m2, "1")
	fmt.Println(m1)
}

func TestFuncMap(t *testing.T) {
	f := func(m map[int]bool) {
		delete(m, 1)
	}
	m := map[int]bool{1: true}
	f(m)
	fmt.Println(m)
}

func TestSyncMap(t *testing.T) {
	sm := sync.Map{}
	sm.Store(1, 1)
	sm.Store("abc", "cba")
	sm.Store(nil, nil)
	fmt.Println(sm.Load(nil))
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Printf("%065b\n", ^uintptr(0))
}

func TestMapExpand(t *testing.T) {
	m := make(map[int]int, 0)
	for i := 0; i < 100; i++ {
		m[i] = i
		fmt.Printf("%p\t%v\n", m, m)
	}
}

func TestMapComplex(t *testing.T) {
	m := make(map[complex64]int)
	m[1+1i] = 100
	fmt.Println(m[1+1i])
	x := complex64(1 + 2i)
	fmt.Println(real(x), imag(x))
	a, b := 1, 2
	x = complex(float32(a), float32(b))
	fmt.Println(x)
	fmt.Println(real(x), imag(x))
}

type s struct {
	A int
}

func TestModify(t *testing.T) {
	m := make(map[int]s)
	for i := 0; i < 10; i++ {
		m[i] = s{A: i}
	}
	fmt.Println(m)
	for k, v := range m {
		v.A += k
	}
	fmt.Println(m)
	for k, v := range m {
		v.A += k
		m[k] = v
	}
	fmt.Println(m)
	fmt.Println(m)
}

func TestRandMap(t *testing.T) {
	m := make(map[int]struct{})
	for i := 0; i < 10; i++ {
		m[i] = struct{}{}
	}
	total := make(map[int]int)
	for i := 0; i < 10000; i++ {
		index := rand.Intn(len(m))
		for k := range m {
			if index == 0 {
				total[k]++
				break
			}
			index--
		}
	}
	fmt.Println(total)
}

var m = map[int]int{}

func TestMapLoop(t *testing.T) {
	initMap()
	go read()
	write()
	time.Sleep(time.Second)
}

func initMap() {
	for i := 0; i < 10; i++ {
		m[10+i] = 10 + i
	}
}

func read() {
	for k, v := range m {
		fmt.Println(k, v)
		time.Sleep(time.Millisecond * 100)
	}
}

func write() {
	for i := 0; i < 10; i++ {
		m[i] = i
		time.Sleep(time.Millisecond * 100)
	}
}

// 切片去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	length := len(arr)
	result := make([]string, 0, length)
	temp := map[string]struct{}{}
	for i := 0; i < length; i++ {
		str := arr[i]
		if _, ok := temp[str]; !ok { // 如果字典中找不到元素，ok=false，!ok为true，就往切片中append元素。
			temp[str] = struct{}{}
			result = append(result, str)
		}
	}
	return result
}

func TestConcurrentAppend(t *testing.T) {
	arr := []string{"hello", "world"}
	for i := 0; i < 100; i++ {
		go func() {
			arr = append(arr, "hello")
			fmt.Println(arr)
			arr = RemoveRepeatedElement(arr)
		}()
	}
	time.Sleep(5 * time.Second)
}

func TestSA(b *testing.T) {
	var locker sync.WaitGroup
	for j := 0; j < 10; j++ {
		arr := []string{"123", "1234", "12345", "12356", "1", "2"}
		for i := 0; i < 30000; i++ {
			locker.Add(1)
			go func() {
				RemoveRepeatedElement(arr)
				locker.Done()
			}()
			go func() {
				arr = append(arr, "123")
			}()
		}
	}
	locker.Wait()
}

func TestSyncSlice(b *testing.T) {
	var locker sync.WaitGroup
	for j := 0; j < 10; j++ {
		arr := newSS("123", "1234", "12345", "12356", "1", "2")
		for i := 0; i < 30000; i++ {
			locker.Add(1)
			go func() {
				arr.RemoveRepeatedElement()
				locker.Done()
			}()
			go func() {
				arr.Append("123")
			}()
		}
	}
	locker.Wait()
}

type SyncSlice struct {
	data []interface{}
	lock *sync.RWMutex
}

func newSS(data ...interface{}) *SyncSlice {
	return &SyncSlice{
		data: data,
		lock: &sync.RWMutex{},
	}
}

func (s *SyncSlice) Append(i interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, i)
}

func (s *SyncSlice) RemoveRepeatedElement() []interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	length := len(s.data)
	result := make([]interface{}, 0, length)
	temp := map[interface{}]struct{}{}
	for _, d := range s.data {
		if _, ok := temp[d]; !ok { // 如果字典中找不到元素，ok=false，!ok为true，就往切片中append元素。
			temp[d] = struct{}{}
			result = append(result, d)
		}
	}
	return result
}
