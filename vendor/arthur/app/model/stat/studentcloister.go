/*
Created on 2018-10-31 13:45:24
author: Auto Generate
*/
package stat

type StudentCloister struct {
	Level        int16   `db:"smallint(5)"` //修道院等级
	NeedExp      int     `db:"int(10)"`     //升级所需
	Consume      int16   `db:"smallint(5)"` //升级消耗钻石
	StudentLimit int16   `db:"smallint(5)"` //学员上限
	Subjects     []int16 `db:"varchar(36)"` //已解锁科目
	InspireTimes int16   `db:"smallint(5)"` //修道院鼓励次数
}
