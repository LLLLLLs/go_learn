/*
Created on 2018-07-03 19:10:13
author: Auto Generate
*/
package stat

import "arthur/app/enum/AttrType"

type HeroTalent struct {
	No      int16         `model:"pk"` //
	AttrTyp AttrType.Type //属性类型：智：1, 勤：2, 忠：3, 英：4
	Star    int16         //星级
	Name    string        //资质名称
}
