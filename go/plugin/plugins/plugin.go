//@author: lls
//@time: 2020/04/26
//@desc:

package main

func main() {}

type IPerson interface {
	Name() string
}

type person struct {
	name string
}

func (p person) Name() string {
	return p.name
}

func NewPerson(name string) IPerson {
	return person{name: name}
}
