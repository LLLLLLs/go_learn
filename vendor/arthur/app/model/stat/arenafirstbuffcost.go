/*
Created on 2018-10-19 16:32:26
author: Auto Generate
*/
package stat

type ArenaFirstBuffCost struct {
	No        int16 `model:"pk"` //buff编号
	BuffTyp   int16 //buff类型，0为攻击，1为血量
	BuffValue int16 //buff的值
	CostType  int16 //花费类型，0为钻石，1为宝珠
	CostValue int16 //宝珠花费
}
