// Time        : 2020/05/27
// Description :

package implement2

import "fmt"

type State string

type Handle func()

type Handles []Handle

func (hs Handles) Handle() Handle {
	return func() {
		for i := range hs {
			hs[i]()
		}
	}
}

type FSM interface {
	State() State
	ToState(s State)
	Register(from, to State, handle ...Handle)
}

// Finite State Machine
type fsm struct {
	state   State
	handles map[State]map[State]Handle // from to handle
}

func NewFSM(state State) FSM {
	return &fsm{
		state:   state,
		handles: make(map[State]map[State]Handle),
	}
}

func (f fsm) State() State {
	return f.state
}

func (f *fsm) ToState(s State) {
	if f.state == s {
		return
	}
	if _, ok := f.handles[f.state]; !ok {
		panic("no handles map")
	}
	handle, ok := f.handles[f.state][s]
	if !ok {
		panic("no handles")
	}
	handle()
	f.state = s
}

func (f *fsm) Register(from, to State, handle ...Handle) {
	if _, ok := f.handles[from]; !ok {
		f.handles[from] = make(map[State]Handle)
	}
	if _, ok := f.handles[from][to]; ok {
		panic(fmt.Sprintf("%s to %s is registerd", from, to))
	}
	f.handles[from][to] = append(Handles{}, handle...).Handle()
}
