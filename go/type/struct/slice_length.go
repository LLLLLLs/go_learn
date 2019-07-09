// Time        : 2019/07/09
// Description : 计算slice占用内存长度

package _struct

type foo struct {
	a int
	b string
}

type bar struct {
	c float64
}

type Aggregation struct {
	field1 []foo
	field2 []bar
}

var agg = Aggregation{}
