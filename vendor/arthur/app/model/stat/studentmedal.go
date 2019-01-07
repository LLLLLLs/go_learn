/*
Created on 2018-10-31 15:07:27
author: Auto Generate
*/
package stat

type StudentMedal struct {
	Id            int16 `model:"pk"` //勋章等级
	Name          string             //勋章名称
	NeedAttr      int                //所需属性
	CloisterExp   int16              //毕业后增加修道院经验
	GoldConsume   int16              //联姻元宝消耗
	NeedItemNo    int16              //联姻消耗道具编号
	NeedItemCount int16              //联姻消耗道具数量
	Award         string             //联姻成功奖励
}
