package HeroProfNo

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Warrior Type = 1 //战士
	Knight  Type = 2 //骑士
	Priest  Type = 3 //教士
	Mage    Type = 4 //法师
)
