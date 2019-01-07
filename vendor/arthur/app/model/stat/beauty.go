/*
Created on 2018-10-19 14:30:32
author: Auto Generate
*/
package stat

type Beauty struct {
	BeautyType    int      `model:"pk"  db:"int(11)"` //对应红颜id
	Name          string   `db:"varchar(255)"`        //红颜名称
	NameFrom      string   `db:"varchar(255)"`        //红颜信息面板，名字边上的小名
	ChatContent   string   `db:"varchar(255)"`        //对话轮播
	BeautyDesc    string   `db:"varchar(255)"`        //红颜描述
	HeroNo        int16    `db:"smallint(11)"`        //对应英雄编号
	Skills        []int    `db:"varchar(255)"`        //拥有家族技能id列表
	Normal        bool     `db:"tinyint(4)"`          //是否是普通家族1是， 0否
	InitCharm     int      `db:"int(255)"`            //初始魅力
	OptionText    string   `db:"varchar(255)"`        //家族解锁途径
	ThankContent  string   `db:"varchar(255)"`        //
	SingleStoryId int      `db:"int(11)"`             //单独邀约剧情ID
	FavorEvent    []string `db:"varchar(255)"`        //名媛解锁前剧情列表
	FriendEvent   string   `db:"varchar(255)"`        //解锁后剧情列表
	ChipEvent     string   `db:"varchar(255)"`        //解锁后碎片剧情列表
	Free          bool     `db:"tinyint(4)"`          //
	FavoriteSet   []int    `db:"varchar(255)"`        //喜欢的物品集合
}
