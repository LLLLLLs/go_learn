// Time        : 2019/07/12
// Description :

package restore_ip_address_93

import (
	"fmt"
	"testing"
)

func TestRestoreIp(t *testing.T) {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("1111"))
	fmt.Println(restoreIpAddresses("00000"))
}

func BenchmarkRestoreIp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restoreIpAddresses("25525511135")
	}
}
