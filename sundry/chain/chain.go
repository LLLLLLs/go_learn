//@author: lls
//@time: 2020/05/19
//@desc:

package chain

import context2 "context"

type Chain interface {
	Handle(c Context) error
}

/*
	A
	|	  B
	|→ → →|	    C
	|	  |→ → →|   Business
	|	  |	    | → → →|
	|	  |← ← ←|
	|← ← ←|
	|
*/
type Context interface {
	context2.Context
	Append(c ...Chain)
	SetValue(key interface{}, value interface{})
	Next() error
}

type context struct {
	context2.Context
	curIndex int
	chains   []Chain
}

func NewContext() Context {
	return &context{
		Context:  context2.Background(),
		curIndex: -1,
		chains:   make([]Chain, 0),
	}
}

func (c *context) Append(chain ...Chain) {
	c.chains = append(c.chains, chain...)
}

func (c *context) SetValue(key, value interface{}) {
	c.Context = context2.WithValue(c.Context, key, value)
}

func (c *context) Next() error {
	c.curIndex++
	if c.curIndex >= len(c.chains) {
		return nil
	}
	return c.chains[c.curIndex].Handle(c)
}
