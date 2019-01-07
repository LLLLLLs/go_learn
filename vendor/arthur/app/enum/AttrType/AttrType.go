package AttrType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Wise     Type = 1 //智慧
	Diligent Type = 2 //勤勉
	Loyalty  Type = 3 //忠诚
	Heroic   Type = 4 //英勇
)

var List = [4]Type{Wise, Diligent, Loyalty, Heroic}
