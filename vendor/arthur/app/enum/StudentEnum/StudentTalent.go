/*
Author      : lls
Time        : 2018/10/31
Description : 
*/

package StudentEnum

type TalentLevel int16

func (t TalentLevel) ToInt16() int16 {
	return int16(t)
}

const (
	I    TalentLevel = 1
	II   TalentLevel = 2
	III  TalentLevel = 3
	IV   TalentLevel = 4
	V    TalentLevel = 5
	VI   TalentLevel = 6
	VII  TalentLevel = 7
	VIII TalentLevel = 8
)
