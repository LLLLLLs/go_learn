//@author: lls
//@time: 2020/05/26
//@desc:

package rdd_test

import (
	. "golearn/sundry/rdd/implement1"
	"testing"
)

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
