//@author: lls
//@time: 2020/05/26
//@desc: 职责驱动设计(responsibility driven design)

package rdd

import (
	"fmt"
	telnet2 "golearn/sundry/rdd/telnet"
)

type State interface {
	SwitchToInitial(t telnet2.Telnet)
	SwitchToMBS(t telnet2.Telnet) // Master + Shell
	SwitchToMBU(t telnet2.Telnet) // Master + UShell
	SwitchToSBS(t telnet2.Telnet) // Slave + Shell
	SwitchToSBU(t telnet2.Telnet) // Slave + UShell
}

type switchToSelf struct{}

func (s switchToSelf) SwitchToInitial(telnet2.Telnet) {}

func (s switchToSelf) SwitchToMBS(telnet2.Telnet) {}

func (s switchToSelf) SwitchToMBU(telnet2.Telnet) {}

func (s switchToSelf) SwitchToSBS(telnet2.Telnet) {}

func (s switchToSelf) SwitchToSBU(telnet2.Telnet) {}

type initial struct{ switchToSelf }

func (i initial) SwitchToMBS(t telnet2.Telnet) {
	t.Login()
}

func (i initial) SwitchToMBU(t telnet2.Telnet) {
	t.Login()
	t.EnterUShell()
}

func (i initial) SwitchToSBS(t telnet2.Telnet) {
	t.Login()
	t.EnterSlave()
}

func (i initial) SwitchToSBU(t telnet2.Telnet) {
	t.Login()
	t.EnterSlave()
	t.EnterUShell()
}

type masterShell struct{ switchToSelf }

func (m masterShell) SwitchToInitial(t telnet2.Telnet) {
	t.Logout()
}

func (m masterShell) SwitchToMBU(t telnet2.Telnet) {
	t.EnterUShell()
}

func (m masterShell) SwitchToSBS(t telnet2.Telnet) {
	t.EnterSlave()
}

func (m masterShell) SwitchToSBU(t telnet2.Telnet) {
	t.EnterSlave()
	t.EnterUShell()
}

type masterUShell struct{ switchToSelf }

func (m masterUShell) SwitchToInitial(t telnet2.Telnet) {
	t.ExitUShell()
	t.Logout()
}

func (m masterUShell) SwitchToMBS(t telnet2.Telnet) {
	t.ExitUShell()
}

func (m masterUShell) SwitchToSBS(t telnet2.Telnet) {
	t.ExitUShell()
	t.EnterSlave()
}

func (m masterUShell) SwitchToSBU(t telnet2.Telnet) {
	t.EnterSlave()
}

type slaveShell struct{ switchToSelf }

func (s slaveShell) SwitchToInitial(t telnet2.Telnet) {
	t.ExitSlave()
	t.Logout()
}

func (s slaveShell) SwitchToMBS(t telnet2.Telnet) {
	t.ExitSlave()
}

func (s slaveShell) SwitchToMBU(t telnet2.Telnet) {
	t.ExitSlave()
	t.EnterUShell()
}

func (s slaveShell) SwitchToSBU(t telnet2.Telnet) {
	t.EnterUShell()
}

type slaveUShell struct{ switchToSelf }

func (s slaveUShell) SwitchToInitial(t telnet2.Telnet) {
	t.ExitUShell()
	t.ExitSlave()
	t.Logout()
}

func (s slaveUShell) SwitchToMBS(t telnet2.Telnet) {
	t.ExitUShell()
	t.ExitSlave()
}

func (s slaveUShell) SwitchToMBU(t telnet2.Telnet) {
	t.ExitSlave()
}

func (s slaveUShell) SwitchToSBS(t telnet2.Telnet) {
	t.ExitUShell()
}

type telnet struct {
	stat State
}

func NewInitTelnet() telnet2.Telnet {
	return &telnet{stat: initial{}}
}

func (t *telnet) Transfer() {
	t.stat.SwitchToMBS(t)
	t.stat = masterShell{}
	fmt.Println("transfer")
}

func (t *telnet) Active() {
	t.stat.SwitchToMBU(t)
	t.stat = masterUShell{}
	fmt.Println("active")
}

func (t *telnet) Inactive() {
	t.stat.SwitchToMBU(t)
	t.stat = masterUShell{}
	fmt.Println("inactive")
}

func (t *telnet) Put() {
	t.stat.SwitchToSBS(t)
	t.stat = slaveShell{}
	fmt.Println("put")
}

func (t *telnet) Get() {
	t.stat.SwitchToSBS(t)
	t.stat = slaveShell{}
	fmt.Println("get")
}

func (t *telnet) DeleteFiles() {
	t.stat.SwitchToSBU(t)
	t.stat = slaveUShell{}
	fmt.Println("delete files")
}

func (t *telnet) Exit() {
	t.stat.SwitchToInitial(t)
	t.stat = initial{}
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

func (t *telnet) State() State {
	return t.stat
}

func (t *telnet) Switch(s State) {
	t.stat = s
}
