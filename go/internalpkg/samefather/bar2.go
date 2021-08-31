// Time        : 2019/11/01
// Description :

package samefather

import "golearn/go/internalpkg/samefather/internal/foo"

// go 特殊文件夹:internal
// 只有同个父目录下的文件夹才有访问权限
// 如当前目录有:go_learn/go/internal,go_learn/go/defer,go_learn/sundry
// 那么defer目录下的文件可调用Get()函数,sundry下的文件无法调用。
// go_learn/go/internal/foo/foo同上
func Get() int {
	foo.Foo()
	return 123
}
