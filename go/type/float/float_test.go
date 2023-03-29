// Time        : 2019/01/10
// Description :

package float

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func BenchmarkAddFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addFloat()
	}
}

func BenchmarkAddZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addZero()
	}
}

// float32二进制组成:23位有效数字+7位指数值+1位小数点方向+1位符号
// 针对123456789进行分析:
// 二进制数为:111 0101 1011 1100 1101 0001 0101‬共27位
// 将小数点往左移动26次得到1.11 0101 1011 1100 1101 0001 0101‬
// 默认省略首位的1，从小数点有点开始数23位取得
//	11 0101 1011 1100 1101 0001 0
// 剩下3位101将被截断，截断规则:4舍6入5取偶(测试得知)，此处101应进位
// 因为最终有效数字为11 0101 1011 1100 1101 0001 1
// 再将二进制转为10进制得111 0101 1011 1100 1101 0001 1000 = 123456792
// 截断规则测试过程：
// 	111 0101 1011 1100 1101 0001 0101(123456789)‬ ==> 111 0101 1011 1100 1101 0001 1000(123456792)
// 	111 0101 1011 1100 1101 0001 1011(123456795)‬ ==> 111 0101 1011 1100 1101 0001 1000(123456792)
// 	111 0101 1011 1100 1101 0001 1100(123456796)‬ ==> 111 0101 1011 1100 1101 0010 0000(123456800)
// 	111 0101 1011 1100 1101 0001 0100(123456788)‬ ==> 111 0101 1011 1100 1101 0001 0000(123456784)
// 	111 0101 1011 1100 1101 0000 1100(123456780)‬ ==> 111 0101 1011 1100 1101 0001 0000(123456784)
func TestPrintFloat(t *testing.T) {
	var f = float32(123456780)
	// 123456792.000000
	fmt.Printf("%f\n", f)
}

// 111 0101 1011 1100 1101 0001 0101‬
// 111 0101 1011 1100 1101 0001 1000
// 111 0101 1011 1100 1101 0001 1011
//
// 110 1011 0111 1001 1010 0011 0011 010 1 1

func TestPrintFloat2(t *testing.T) {
	fmt.Printf("%.64f\n", 0.2)
	fmt.Printf("%.64f\n", 0.1)
	fmt.Printf("%.64f\n", 0.2+0.1)
	fmt.Printf("%.64f\n", 0.3)
	a := 0.1
	b := 0.2
	fmt.Println(a + b)
	fmt.Printf("%.64f\n", a+b)
	fmt.Printf("%.64f\n", 0.5)
	fmt.Printf("%b\n", 0.2)
	fmt.Printf("%b\n", 0.1)
	fmt.Printf("%b\n", 0.5)
}

func TestPrintFloat3(t *testing.T) {
	a := 0.1
	b := 0.2
	fmt.Printf("%.54f\n", a+b)
	fmt.Printf("%.54f\n", 0.1+0.2)
	fmt.Printf("%.54f\n", (a+b)*1e10)
	fmt.Printf("%.54f\n", (0.1+0.2)*1e1)
	fmt.Printf("%.54f\n", a*1e1+b*1e1)
}

func TestFloat64ToInt64(t *testing.T) {
	a := int64(math.MaxInt64)
	b := int64(math.MaxInt64)
	c := int64(float64(a) * float64(b))
	fmt.Println(c)
	d := int64(1)
	fmt.Println(float64(a) + float64(d))

	ast := assert.New(t)
	e := float64ToInt64(float64(math.MaxInt64) + float64(100))
	ast.Equal(int64(math.MaxInt64), e)
	f := float64ToInt64(float64(math.MaxInt64) - float64(100))
	ast.Equal(int64(math.MaxInt64-100), f)
}

func TestTruncate(t *testing.T) {
	a := 101
	b := 0.8
	c := int64(200)
	d := int64(float64(a)*b) * c
	fmt.Println(d)
	e := int64(math.Floor(float64(101)*0.8)) * 200
	fmt.Println(e)
}

func TestFloatConvert(t *testing.T) {
	a := int64(1596059364005777408)
	b := float64(1596059364005777408)
	c := int64(b)
	fmt.Println(a, b, c)
}
