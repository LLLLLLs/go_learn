// Time        : 2019/08/23
// Description :

package goroutine

import "testing"

func TestWithGoroutine(t *testing.T) {
	sumWithGoroutine()
}

func TestWithoutGoroutine(t *testing.T) {
	sumWithoutGoroutine()
}
