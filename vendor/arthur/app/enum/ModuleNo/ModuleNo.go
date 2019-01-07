package ModuleNo

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Hero         Type = 1
	Diplomacy    Type = 2
	Backpack     Type = 3
	Store        Type = 4
	Visit        Type = 5
	Phase        Type = 6
	Beauty       Type = 7
	Mail         Type = 8
	Rank         Type = 9
	RoleInfo     Type = 10
	SOK          Type = 11
	Arena        Type = 12
	WishPool     Type = 13
	Alliance     Type = 14
	AllianceTask Type = 15
	Chat         Type = 16
	Student      Type = 17
	Marriage     Type = 18
)
