// Time        : 2019/07/03
// Description :

package count_II_52

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	result := totalNQueens(4)
	fmt.Println(result)
	result = totalNQueens(8)
	fmt.Println(result)
}
