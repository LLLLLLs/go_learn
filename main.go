package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var stepMap = map[int][][]int{
	1: {{1}},
	2: {{1, 1}, {2}},
}

type kv struct {
	k int
	v string
}

func sliceSeparate() {
	rand.Seed(time.Now().Unix())
	kvs := make([]kv, 100)
	inc := 0
	for i := 0; i < 100; i++ {
		if rand.Intn(100)+1 < 30 {
			inc++
		}
		kv := kv{
			k: inc,
			v: "string" + strconv.Itoa(i),
		}
		kvs[i] = kv
	}

	kvss := make([][]kv, 0)
	k := kvs[0].k
	kvss = append(kvss, make([]kv, 0))
	var index1 = 0
	for i := 0; i < 100; i++ {
		if kvs[i].k != k {
			index1++
			k++
			kvss = append(kvss, make([]kv, 0))

		}
		kvss[index1] = append(kvss[index1], kvs[i])
	}

	fmt.Println(kvs)
	for i := range kvss {
		fmt.Println(kvss[i])
	}
}

type TestStruct struct {
	A int
	B int
}

func (ts *TestStruct) foo() {
	fmt.Println("this is a foo", ts.A)
}

var (
	xx         = []int{1, 2, 3, 4, 5}
	yy         = []int{2, 3, 4, 5, 6, 7}
	zz         = []int{3, 4, 5, 6, 7, 8, 9}
	ii, jj, kk = 0, 0, 0
)

func Random() int {
	ii++
	jj++
	kk++
	if ii == len(xx)-1 {
		ii = 0
	}
	if jj == len(yy)-1 {
		jj = 0
	}
	if kk == len(zz)-1 {
		kk = 0
	}
	return xx[ii] * yy[jj] * zz[kk]
}

func xxx(f func()) {
	f()
}

func trimUnUsedFunc() {
	a1()
	bitCalc()
	bitTest()
	bubbleSort([]int{1, 2, 3})
	calcBlood([]int64{1, 2, 3, 4}, 5)
	calcHits()
	calGrowSpeed()
	callJosephus()
	calPower(1, 2, 3)
	ceil(111)
	chanWithContext()
	compQAndBSort()
	_, _ = Contain(1, []int{1, 2, 3})
	deferCall()
	divide2([]int{1, 2, 3})
	formatTime()
	genRandTestStructList()
	hex()
	king([][2]int{{1, 2}}, 1)
	letterChange("abc")
	listToString([]int{1, 2, 3})
	log("log")
	marshal()
	marshalTest()
	parseInt()
	parseStudent()
	printSliceWithGoroutine()
	var ab = &AB{}
	printX(ab)
	printY(ab)
	print1(111)
	qsort1([]int{1, 2, 4, 3, 5})
	Random()
	recursiveWithTail(10, 0, make(chan int))
	returnInt(111)
	round(0.111)
	selectDrop()
	sliceSeparate()
	sortTwice()
	step(2)
	stepFromBottom(2)
	stringToList("123", 10)
	test()
	test1([]int{0})
	testClosure()
	testCondition()
	testDeferInFor()
	testGbq()
	testGenSlice()
	testGoroutine(make(chan int))
	testGoQsort()
	testInterface()
	testList([]int{0})
	testLock("lock", 1)
	testLog()
	testMap(map[string]string{})
	testMapRand()
	testQSort()
	testRefactor1()
	testRefactor2()
	testReflect()
	testRtnSlice()
	testRtn()
	testSelect()
	testSlice([]int{1})
	testSliceAppend()
	testSortStr("")
	testXXX(1)
	ticker()
	xxx(func() {})
}

type locks struct {
	lks map[string]*sync.Mutex
	mux *sync.Mutex
}

var lks = locks{
	lks: make(map[string]*sync.Mutex),
	mux: &sync.Mutex{},
}

func (l *locks) getLock(ls string) *sync.Mutex {
	l.mux.Lock()
	defer l.mux.Unlock()
	lk, ok := l.lks[ls]
	if !ok {
		lk = &sync.Mutex{}
		l.lks[ls] = lk
	}
	return lk
}

func selectDrop() {
	const cap1 = 5
	ch := make(chan string, cap1)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
			//runtime.Gosched()
		default:
			fmt.Println("manager : drop")
		}
	}
	close(ch)
	fmt.Println(runtime.NumCPU())
}

func main() {
	fmt.Println(int64(12345678901) * int64(1234567890))
}

func chanWithContext() {
	ch := make(chan int)
	ticker := time.NewTicker(time.Second)
	count := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	var loop = true
	for loop {
		select {
		case recv := <-ch:
			fmt.Println("chanel close,recv:", recv)
			loop = false
			break
		case <-ticker.C:
			fmt.Println("one second pass")
			count--
			if count == 0 {
				close(ch)
			}
		case <-ctx.Done():
			fmt.Println("ctx time out")
			loop = false
		}
	}
	fmt.Println("recv:", <-ctx.Done())
}

// 查看slice扩展过程
func testSliceAppend() {
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("len:%d\t cap:%d\t ptr:%p\n", len(list), cap(list), list)
		list = append(list, i)
	}
}

func round(f float64) int {
	if f < 0 {
		return int(f - 0.5)
	}
	return int(f + 0.5)
}

// golang condition 使用测试
func testCondition() {
	var count = 4
	ch := make(chan struct{}, 5)

	// 新建 cond
	var l sync.Mutex
	cond := sync.NewCond(&l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			// 条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("得到count的值%d goroutine%d\n", count, i)
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}

			fmt.Printf("goroutine%d 执行结束\n", i)
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Millisecond * 20)
	// 锁定一下，我要改变 count 的值
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}
}

type person struct {
	no int
}

// 解约瑟夫环
func callJosephus() {
	var (
		persons = make([]person, 30)
		index   = 2
		end     = 10
	)
	for i := range persons {
		person := person{
			no: i + 1,
		}
		persons[i] = person
	}
	josephus(persons, index, end)
	josephusImprove(persons, index, end)
}

// list假定为一个环(首位相接) index为死亡序号 end为结束人数
func josephus(persons []person, index, end int) {
	cpy := make([]person, len(persons))
	copy(cpy, persons)
	var count = 1
	for currentIndex := 0; end < len(cpy); {
		if count == index {
			count = 1
			fmt.Printf("%d ", cpy[currentIndex].no)
			cpy = append(cpy[:currentIndex], cpy[currentIndex+1:]...)
			continue
		}
		count++
		currentIndex++
		if currentIndex >= len(cpy) {
			currentIndex = 0
		}
	}
	fmt.Printf("\n")
	fmt.Printf("%v\n", cpy)
}

// list假定为一个环(首位相接) index为死亡序号 end为结束人数 改良版
func josephusImprove(persons []person, index, end int) {
	cpy := make([]person, len(persons))
	copy(cpy, persons)
	currentIndex := 0
	for end < len(cpy) {
		currentIndex = currentIndex + (index - 1)
		if currentIndex >= len(cpy) {
			currentIndex %= len(cpy)
		}
		fmt.Printf("%d ", cpy[currentIndex].no)
		cpy = append(cpy[:currentIndex], cpy[currentIndex+1:]...)
	}
	fmt.Printf("\n")
	fmt.Printf("%v\n", cpy)
}

// 交替打印slice
func printSliceWithGoroutine() {
	var (
		arrA   = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
		arrB   = []string{"a", "b", "c", "d", "e", "f"}
		ok     = make(chan bool, 1)
		num    = make(chan bool, 1)
		letter = make(chan bool, 1)
	)
	go printSlice(arrA, num, letter, ok)
	go printSlice(arrB, letter, num, ok)
	num <- true
	<-ok
	var ctn = true
	for ctn {
		select {
		case <-ok:
			ctn = false
		case <-num:
			letter <- true
		case <-letter:
			num <- true
		}
	}
}

func printSlice(slice interface{}, in, out, ok chan bool) {
	value := reflect.ValueOf(slice)
	for i := 0; i < value.Len(); i++ {
		<-in
		fmt.Println(value.Index(i).Interface())
		out <- true
	}
	ok <- true
}

// 倒计时
func ticker() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("di da")
		}
	}
}

// 内存锁
func testLock(ls string, i int) {
	lk := lks.getLock(ls)
	lk.Lock()
	fmt.Printf("%d got lock %p\n", i, lk)
	defer lk.Unlock()
	time.Sleep(time.Second * 5)
	fmt.Printf("%d exit\n", i)
}

func testGoroutine(q chan int) {
	exit := make(chan int)
	go func() {
		close(exit)
		for {
			if true {
				println("Looping!") //Second
			}
		}
	}()
	<-exit
	println("Am I printed?")
	q <- 1
}

// 反射
func testReflect() {
	i := testInterface()
	value := reflect.ValueOf(i)
	elem := value.Elem()
	fmt.Println(elem.FieldByName("A"))
	fmt.Println(elem.FieldByNameFunc(func(s string) bool {
		return s == "A"
	}))
	fn := reflect.ValueOf((*TestStruct).foo)
	fn.Call([]reflect.Value{value})
}

func testInterface() interface{} {
	return &TestStruct{1, 2}
}

func testRefactor1() {
	var Refactor = 1
	_ = Refactor
}

func testRefactor2() {
	var refactor = 2
	_ = refactor
}

func testDeferInFor() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("222")
	}()
	panic("111")
}

// select测试
func testSelect() {
	var x = make(chan int)
	go input(x)
	for {
		select {
		case <-x:
			fmt.Println(1111)
		}
		fmt.Println(222)
	}
}

func input(x chan int) {
	for {
		time.Sleep(time.Second * 2)
		x <- 1
	}
}

// 位测试
func bitTest() {
	var x uint8 = 1<<1 | 1<<5
	for i := uint8(0); i < 8; i++ {
		fmt.Println(x >> 1)
		fmt.Printf("%08b\n", x)
	}
}

// 位运算
func bitCalc() {
	var (
		x = uint(127)
		y = uint(4)
	)
	z := x &^ y
	fmt.Printf("%08b", z)
}

// 二次排序
func sortTwice() {
	list := genRandTestStructList()
	sort.Slice(list, func(i, j int) bool {
		return list[i].A > list[j].A || (list[i].A == list[j].A && list[i].B > list[j].B)
	})
	fmt.Println(list)
}

// 返回一个ts slice
func genRandTestStructList() []TestStruct {
	list := make([]TestStruct, 100)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		ts := TestStruct{
			A: rand.Intn(100),
			B: rand.Intn(100),
		}
		list[i] = ts
	}
	return list
}

// 计算学员自然出坑所需初始速度
func calGrowSpeed() {
	calPower(1, -1, -76800)
}

// 解一元二次方程
func calPower(a, b, c float64) {
	delta := b*b - 4*a*c
	if delta <= 0 {
		panic("wrong input")
	}
	x1 := (-b + math.Sqrt(float64(delta))) / (2 * a)
	x2 := (-b - math.Sqrt(float64(delta))) / (2 * a)
	fmt.Printf("x1 = %.2f , x2 = %.2f\n", x1, x2)
}

func testMapRand() {
	var data = map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	var count = make(map[int]int)
	var out = 5
	for i := 0; i < 1000; i++ {
		var index = 1
		for k := range data {
			if index == out {
				count[k]++
				break
			}
			index++
		}
	}
	fmt.Printf("随机选择第%d个:%+v", out, count)
}

func recursiveWithTail(num, mid int, result chan int) {
	mid += num
	if num == 1 {
		result <- mid
	}
	go recursiveWithTail(num-1, mid, result)
}

func testLog() *string {
	s := "111"
	sptr := &s
	defer elapsed1(sptr)()
	time.Sleep(time.Millisecond * 10)
	return sptr
}

func elapsed1(n *string) func() {
	start := time.Now()
	fmt.Printf("enter\n")
	return func() {
		*n = fmt.Sprintf("%s", time.Since(start))
	}
}

func log(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() { fmt.Printf("exit %s (%s)", msg, time.Since(start)) }
}

type A interface {
	x(param int)
}

type B interface {
	y(param int)
}

type AB struct {
}

func (ab *AB) x(param int) {
	fmt.Printf("%p", ab)
	fmt.Println(param)
}

func (ab *AB) y(param int) {
	fmt.Printf("%p", ab)
	fmt.Println(param)
}

func printX(a A) {
	fmt.Printf("%p", a)
	a.x(2)
}

func printY(b B) {
	fmt.Printf("%p", b)
	b.y(3)
}

func testSortStr(str string) {
	arr := []rune(str)
	//arr := strings.Split(str, "")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(func(arr []rune) []string {
		ss := make([]string, len(arr))
		for i := range arr {
			ss[i] = string(arr[i])
		}
		return ss
	}(arr))
}

func testRtn() (ts TestStruct) {
	return ts
}

// 数组初始化方式
func testGenSlice() {
	num := 1000000
	start := time.Now().UnixNano()
	s1 := make([]int, 0)
	for i := 0; i < num; i++ {
		s1 = append(s1, i)
	}
	end := time.Now().UnixNano()
	time1 := end - start
	s2 := make([]int, num)
	for i := 0; i < num; i++ {
		s2[i] = i
	}
	time2 := time.Now().UnixNano() - end
	fmt.Println(time1, "VS", time2)
}

// 闭包
func testClosure() {
	jpg := suffix(".jpg")
	txt := suffix(".txt")
	fmt.Println(jpg("hello.jpg"))
	fmt.Println(txt("hello.txt"))
	fmt.Println(jpg("hello.txt"))
}

func suffix(suf string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suf) {
			return name
		}
		return name + suf
	}
}

func testXXX(j int) func(m int) {
	i := 10
	return func(m int) {
		fmt.Println(j, i, m)
	}
}

func testMap(m map[string]string) {
	mm := m
	delete(mm, "A")
	fmt.Println("-----", mm, m)
}

func testRtnSlice() (s []int) {
	s = append(s, 1)
	return
}

func ceil(x float64) int64 {
	return int64(math.Ceil(x))
}

func parseInt() {
	i64, err := strconv.ParseInt("8pivq1ayviaq", 10, 64)
	fmt.Println(err, i64)
}

func formatTime() {
	fmt.Println(time.Unix(1538107572, 0))
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	targetType := reflect.TypeOf(target)
	switch targetType.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

// 根据伤害计算英雄血量,并返回英雄是否全部阵亡
func calcBlood(bloodList []int64, bloodConsume int64) bool {
	var (
		death = 0
	)
	for i := range bloodList {
		if bloodList[i] == 0 {
			death++
		}
	}
	for death != len(bloodList) && bloodConsume != 0 { // 伤害量为0或全部死亡跳出循环
		damage := bloodConsume / int64(len(bloodList)-death)
		bloodConsume = 0
		for i := range bloodList {
			if bloodList[i] == 0 {
				continue
			}
			if bloodList[i]-damage < 0 {
				bloodConsume += damage - bloodList[i]
				bloodList[i] = 0
				death++
				continue
			}
			bloodList[i] -= damage
		}
	}
	return death == len(bloodList)
}

func listToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func marshalTest() {
	str := "[1,2,3,4]"
	var list interface{}
	list = make([]int64, 0)
	err := json.Unmarshal([]byte(str), &list)
	fmt.Println(err)
	fmt.Println(list)
}

func stringToList(str string, base int) interface{} {
	switch base {
	case 16:
		list := make([]int16, 0)
		err := json.Unmarshal([]byte(str), &list)
		fmt.Println(err)
		return list
	case 32:
		list := make([]int, 0)
		err := json.Unmarshal([]byte(str), &list)
		fmt.Println(err)
		return list
	case 64:
		list := make([]int64, 0)
		err := json.Unmarshal([]byte(str), &list)
		fmt.Println(err)
		return list
	default:
		return nil
	}
}

func marshal() {
	l := []int16{1, 2, 3}
	bytes, _ := json.Marshal(l)
	fmt.Println(string(bytes))
	ll := make([]int, 0)
	err := json.Unmarshal(bytes, &ll)
	fmt.Println(err)
	fmt.Println(ll)
}

func hex() {
	base := 'f'
	list := []rune{base}
	for i := 0; i < 15; i++ {
		i64, err := strconv.ParseInt(string(list), 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("hex:%s,\tlen:%d;\tdec:%d,\tlen:%d", string(list), len(list), i64, len([]rune(strconv.FormatInt(i64, 10)))))
		list = append(list, base)
	}
}

func testQSort() {
	rand.Seed(time.Now().Unix())
	var n = 1000000
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(10000)
	}
	start1 := Now()
	betterQSort(list, 0, len(list)-1)
	qtime := Now() - start1
	//fmt.Println(list)
	fmt.Println("time:", qtime)
}

func testGbq() {
	rand.Seed(time.Now().Unix())
	var n = 1000000
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(10000)
	}
	start2 := Now()
	lock.Add(1)
	gbq(list, 0, len(list)-1)
	lock.Wait()
	bqtime := Now() - start2
	fmt.Println("time:", bqtime)
}

func testGoQsort() {
	rand.Seed(time.Now().Unix())
	var n = 100000
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(10000)
	}
	start2 := Now()
	lock.Add(1)
	goQSort(list, 0, len(list)-1)
	lock.Wait()
	bqtime := Now() - start2
	fmt.Println("time:", bqtime)
}

func testSlice(list []int) {
	list[1] = 5
	list = append(list, list...)
	list[3] = 6
	fmt.Println(list)
}

func calcHits() {
	hitPerSecond := 1000000
	total := hitPerSecond * 60 * 60 * 24 * 365
	fmt.Println(total, "length:", len(strconv.Itoa(total)))
}

func compQAndBSort() {
	rand.Seed(time.Now().Unix())
	var n = 1000000
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(10000)
	}
	start1 := Now()
	qs := qsort(list)
	_ = qs
	qtime := Now() - start1
	//time.Sleep(time.Second)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(10000)
	}
	start2 := Now()
	lock.Add(1)
	goQSort(list, 0, len(list)-1)
	lock.Wait()
	bqtime := Now() - start2
	//fmt.Println(list)
	fmt.Printf("quick sort cost\t: %dms\ngo quick sort cost\t: %dms\n", qtime, bqtime)
	//fmt.Println(qs)
}

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func qsort(qlist []int) []int {

	if len(qlist) <= 1 {
		return qlist
	}
	list := make([]int, len(qlist))
	copy(list, qlist)
	var (
		i, j, k = 0, len(list) - 1, 0
	)
	for i < j {
		for i < j && list[j] >= list[k] {
			j--
		}
		list[j], list[k] = list[k], list[j]
		k = j
		for i < j && list[i] <= list[k] {
			i++
		}
		list[i], list[k] = list[k], list[i]
		k = i
	}
	left := qsort(list[:k])
	right := qsort(list[k+1:])
	rtn := append(left, list[k])
	rtn = append(rtn, right...)
	return rtn
}

func qsort1(list []int) {
	var (
		i, j, k = 0, len(list) - 1, 0
	)
	for {
		if i == j {
			break
		}
		for ; j != i; j-- {
			if list[j] < list[k] {
				list[j], list[k] = list[k], list[j]
				k = j
				break
			}
		}
		if j == 0 {
			break
		}
		for ; i != j; i++ {
			if list[i] > list[k] {
				list[i], list[k] = list[k], list[i]
				k = i
				break
			}
		}
	}
	qsort(list[:k])
	qsort(list[k+1:])

	return
}

func betterQSort(list []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	key := list[low]
	for j > i {
		for j > i && list[j] >= key {
			j--
		}
		list[i] = list[j]
		for i < j && list[i] <= key {
			i++
		}
		list[j] = list[i]
	}
	list[j] = key
	betterQSort(list, low, j-1)
	betterQSort(list, j+1, high)
}

var lock sync.WaitGroup

func gbq(list []int, low, high int) {
	defer lock.Done()
	if low >= high {
		return
	}
	i, j := low, high
	key := list[low]
	for j > i {
		for j > i && list[j] >= key {
			j--
		}
		list[i] = list[j]
		for i < j && list[i] <= key {
			i++
		}
		list[j] = list[i]
	}
	list[j] = key
	lock.Add(2)
	go gbq(list, low, j-1)
	go gbq(list, j+1, high)
}

func goQSort(num []int, low, high int) {
	defer lock.Done()

	if low >= high {
		return
	}

	i, j := low, high
	key := num[low]
	for i < j {
		for j > i && num[j] >= key {
			j--
		}
		num[i] = num[j]

		for i < j && num[i] < key {
			i++
		}
		num[j] = num[i]
	}
	num[j] = key

	lock.Add(2)
	go goQSort(num, low, i-1)
	go goQSort(num, i+1, high)
}

func bubbleSort(blist []int) []int {
	list := make([]int, len(blist))
	copy(list, blist)
	for i := len(list) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func king(gold [][2]int, person int) int {
	if person <= 0 {
		return 0
	}
	length := len(gold)
	var flag = true
	if person-gold[length-1][1] < 0 {
		flag = false
	}
	if length == 1 {
		if !flag {
			return 0
		}
		return gold[0][0]
	}
	var value1, value2 = 0, 0
	value1 = king(gold[:length-1], person)
	if flag {
		value2 = king(gold[:length-1], person-gold[length-1][1]) + gold[length-1][0]
	}
	if value1 < value2 {
		return value2
	}
	return value1
}

func stepFromBottom(n int) [][]int {
	resultn1 := [][]int{{1, 1}, {2}}
	resultn2 := [][]int{{1}}
	for i := 3; i <= n; i++ {
		result := make([][]int, 0)
		for _, v := range resultn1 {
			single := make([]int, 0)
			single = append(single, v...)
			single = append(single, 1)
			result = append(result, single)
		}
		for _, v := range resultn2 {
			single := make([]int, 0)
			single = append(single, v...)
			single = append(single, 2)
			result = append(result, single)
		}
		resultn2 = resultn1
		resultn1 = result
	}
	return resultn1
}

func step(n int) [][]int {
	if v, ok := stepMap[n]; ok {
		return v
	}
	//if n == 1 {
	//	return [][]int{{1}}
	//}
	//if n == 2 {
	//	return [][]int{{1, 1}, {2}}
	//}
	result := make([][]int, 0)
	step1, step2 := step(n-1), step(n-2)
	for _, v := range step1 {
		single := make([]int, 0)
		single = append(single, v...)
		single = append(single, 1)
		result = append(result, single)
	}
	for _, v := range step2 {
		single := make([]int, 0)
		single = append(single, v...)
		single = append(single, 2)
		result = append(result, single)
	}
	stepMap[n] = result
	return result
}

func divide2(list []int) []int {
	if len(list) == 1 {
		return list
	}
	index := len(list) / 2
	l1, l2 := list[:index], list[index:]
	l1, l2 = divide2(l1), divide2(l2)
	return merge(l1, l2)
}

func merge(list1, list2 []int) []int {
	rtn := make([]int, 0)
	var i, j, k = 0, 0, 0
	for ; i < len(list1)+len(list2); i++ {
		if j == len(list1) || k == len(list2) {
			break
		}
		if list1[j] < list2[k] {
			rtn = append(rtn, list1[j])
			j++
		} else {
			rtn = append(rtn, list2[k])
			k++
		}
	}
	if j == len(list1) {
		rtn = append(rtn, list2[k:]...)
	} else {
		rtn = append(rtn, list1[j:]...)
	}
	return rtn
}

func letterChange(str string) bool {
	runes := []rune(str)
	for i, v := range runes {
		if v >= 'a' && v <= 'z' {
			if i == 0 || i == len(runes)-1 || (runes[i-1] != '+' || runes[i+1] != '+') {
				return false
			}
		}
	}
	return true
}

func print1(num int) {
	fmt.Println(num)
}

func returnInt(num int) int {
	return num
}

func testList(list []int) {
	for i, j := range list {
		if i == j {
			list = append(list[:i], list[i+1:]...)
		}
	}
}

func test1(ss []int) {
	for _, s := range ss {
		s++
	}
	ss = append(ss, 1, 2, 3, 4, 5, 6)
	fmt.Println(ss)
}

func a1() (a int, b int) {
	c := 1
	return c, b
}

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	ss := []student{
		{Name: "zha", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range ss {
		m[stu.Name] = &stu
	}
	for _, stu := range m {
		fmt.Println(stu)
	}
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func test() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
