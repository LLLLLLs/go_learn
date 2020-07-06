//@author: lls
//@time: 2020/05/26
//@desc:

package implement2

import (
	"fmt"
	telnet2 "golearn/sundry/rdd/telnet"
)

type telnet struct {
	fsm FSM
}

func NewInitTelnet() telnet2.Telnet {
	t := &telnet{}
	fsm := NewFSM(Initial)
	fsm.Register(Initial, MasterShell, t.Login)
	fsm.Register(Initial, MasterUShell, t.Login, t.EnterUShell)
	fsm.Register(Initial, SlaveShell, t.Login, t.EnterSlave)
	fsm.Register(Initial, SlaveUShell, t.Login, t.EnterSlave, t.EnterUShell)

	fsm.Register(MasterShell, Initial, t.Logout)
	fsm.Register(MasterShell, MasterUShell, t.EnterUShell)
	fsm.Register(MasterShell, SlaveShell, t.EnterSlave)
	fsm.Register(MasterShell, SlaveUShell, t.EnterSlave, t.EnterUShell)

	fsm.Register(MasterUShell, Initial, t.ExitUShell, t.Logout)
	fsm.Register(MasterUShell, MasterShell, t.ExitUShell)
	fsm.Register(MasterUShell, SlaveShell, t.ExitUShell, t.EnterSlave)
	fsm.Register(MasterUShell, SlaveUShell, t.EnterSlave)

	fsm.Register(SlaveShell, Initial, t.ExitSlave, t.Logout)
	fsm.Register(SlaveShell, MasterShell, t.ExitSlave)
	fsm.Register(SlaveShell, MasterUShell, t.ExitUShell, t.EnterUShell)
	fsm.Register(SlaveShell, SlaveUShell, t.EnterUShell)

	fsm.Register(SlaveUShell, Initial, t.ExitUShell, t.ExitSlave, t.Logout)
	fsm.Register(SlaveUShell, MasterShell, t.ExitSlave, t.ExitUShell)
	fsm.Register(SlaveUShell, MasterUShell, t.ExitSlave)
	fsm.Register(SlaveUShell, SlaveShell, t.ExitUShell)
	t.fsm = fsm
	return t
}

func (t *telnet) Transfer() {
	t.fsm.ToState(MasterShell)
	fmt.Println("transfer")
}

func (t *telnet) Active() {
	t.fsm.ToState(MasterUShell)
	fmt.Println("active")
}

func (t *telnet) Inactive() {
	t.fsm.ToState(MasterUShell)
	fmt.Println("inactive")
}

func (t *telnet) Put() {
	t.fsm.ToState(SlaveShell)
	fmt.Println("put")
}

func (t *telnet) Get() {
	t.fsm.ToState(SlaveShell)
	fmt.Println("get")
}

func (t *telnet) DeleteFiles() {
	t.fsm.ToState(SlaveUShell)
	fmt.Println("delete files")
}

func (t *telnet) Exit() {
	t.fsm.ToState(Initial)
	fmt.Println("exit")
}

func (t *telnet) Login() {
	fmt.Println("login")
}

func (t *telnet) Logout() {
	fmt.Println("logout")
}

func (t *telnet) EnterSlave() {
	fmt.Println("enter slave")
}

func (t *telnet) ExitSlave() {
	fmt.Println("exit slave")
}

func (t *telnet) EnterUShell() {
	fmt.Println("enter u shell")
}

func (t *telnet) ExitUShell() {
	fmt.Println("exit u shell")
}
