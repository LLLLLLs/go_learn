// Time        : 2019/07/09
// Description :

package simplify_path_71

import (
	"fmt"
	"testing"
)

func TestSimplify(t *testing.T) {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
