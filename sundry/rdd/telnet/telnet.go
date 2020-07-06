// Time        : 2020/05/27
// Description :

package telnet

type Telnet interface {
	HandleCommand
	StateChangeCommand
}

type HandleCommand interface {
	Transfer()
	Active()
	Inactive()
	Put()
	Get()
	DeleteFiles()
	Exit()
}

type StateChangeCommand interface {
	Login()
	Logout()
	EnterSlave()
	ExitSlave()
	EnterUShell()
	ExitUShell()
}
