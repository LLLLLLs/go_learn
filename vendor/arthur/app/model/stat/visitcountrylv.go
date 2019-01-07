/*
Created on 2018-10-18 16:10:54
author: Auto Generate
*/
package stat

type VisitCountryLv struct {
	Lv               int     `model:"pk"` //国家等级
	NeedExp          int     //经验需求
	NeedAttr         int64   //属性需求
	ResourceAward    int64   //基础资源奖励
	TalentAdd        float64 //英雄天赋加成
	AwardCoefficient float64 //奖励参数
	AttrAward        int     //英雄属性奖励
	NeedSilver       int64   //金币需求
	NeedFood         int64   //粮食需求
	AttrLimit        float64 //属性比值上限
}
