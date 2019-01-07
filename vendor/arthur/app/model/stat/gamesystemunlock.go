/*
Created on 2018-12-14 09:20:42
author: Auto Generate
*/
package stat

import "arthur/app/enum/UnlockType"

type GameSystemUnlock struct {
	No    int16 `model:"pk"  db:"int(6)"` //编号
	Type  UnlockType.Type `db:"int(6)"`             //1通关关卡章数，2达到荣誉等级，3英雄数量，4名媛数量，5子嗣数量
	Value int `db:"int(6)"`             //具体值
}
