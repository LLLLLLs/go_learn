/*
Author      : lls
Time        : 2018/10/31
Description : 
*/

package StudentEnum

type Stat int16

func (stat Stat) ToInt16() int16 {
	return int16(stat)
}

var (
	Growing      Stat = 1 // 成长中
	MidExam      Stat = 2 // 期中考
	FinalExam    Stat = 3 // 期末考
	WaitGraduate Stat = 4 // 等待授勋
	Graduated    Stat = 5 // 已毕业
)

type StatForClient int16

func (s StatForClient) ToInt16() int16 {
	return int16(s)
}

type clientStat struct {
	Escape       int16 // 逃学状态
	Growing      int16 // 成长状态
	WaitExam     int16 // 等待考试
	Examing      int16 // 考试中
	ExamFail     int16 // 考试失败
	WaitGraduate int16 // 等待毕业
	ExamReset    int16 // 考试失败并且使用过道具
}

var ClientStat = clientStat{
	Escape:       1,
	Growing:      2,
	WaitExam:     3,
	Examing:      4,
	ExamFail:     5,
	WaitGraduate: 6,
	ExamReset:    7,
}
