/*
Created on 2018-12-06 11:09:01
author: Auto Generate
*/
package stat

type StudentEncourage struct {
	No     int `model:"pk"  db:"int(6)"` //有效在坑人数
	Effect int `db:"int(6)"`             //鼓励效果
}
