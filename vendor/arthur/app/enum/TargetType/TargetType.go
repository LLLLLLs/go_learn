package TargetType

import "arthur/app/info/errors"

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

func (t Type) ToString() string {
	switch t {
	case Role:
		return "Role"
	case Hero:
		return "Hero"
	case Beauty:
		return "Beauty"
	case Item:
		return "Item"
	default:
		panic(errors.ErrType)
	}
}

const (
	Role   Type = 1
	Hero   Type = 2
	Beauty Type = 3
	Item   Type = 4
)

var (
	List = []Type{
		Role,
		Hero,
		Beauty,
		Item,
	}
)
