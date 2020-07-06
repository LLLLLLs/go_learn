// Time        : 2020/05/27
// Description :

package implement2

import "testing"

func TestUpgrade(t *testing.T) {
	telnet := NewInitTelnet()
	telnet.Inactive()
	telnet.Get()
	telnet.Transfer()
	telnet.Put()
	telnet.DeleteFiles()
	telnet.Active()
	telnet.Exit()
}
