//@author: lls
//@time: 2020/05/19
//@desc: godoc测试,godoc会生成GOROOT和GOPATH下的所有文档,因此需要将目标目录加到GOPATH中,
//并且go v1.12.0后不再内置该命令,需要从golang.org/x/tools/cmd/godoc中获取。

package queue

type Queue []interface{}

// Push 队列末尾插入一个元素
func (p *Queue) Push(v interface{}) {
	*p = append(*p, v)
}

// Pop 获得队列头
func (p *Queue) Pop() interface{} {
	head := (*p)[0]
	*p = (*p)[1:]
	return head
}

// IsEmpty 判断队列是否为空
func (p *Queue) IsEmpty() bool {
	return len(*p) == 0
}
