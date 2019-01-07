/*
Created on 2018-07-09 09:32:20
author: Auto Generate
*/
package stat

import "arthur/app/enum/HeroProfNo"

type HeroProf struct {
	No        HeroProfNo.Type `model:"pk"` //
	Coeff     float64         //职业系数
	AttackTyp int16           //攻击类型
	AttackCoeff float64		//攻击系数
	BloodCoeff float64		//血量系数
}
