/*
Created on 2018-11-22 13:52:27
author: Auto Generate
*/
package stat

type StoreItem struct {
	Id       int16 `model:"pk"  db:"smallint(6)"` //主键，无意义
	ItemNo   int16 `db:"smallint(6)"`             //道具编号
	ItemType int16 `db:"smallint(6)"`             //道具类型(1道具，2值，3英雄，4红颜)
	ItemNum  int16 `db:"smallint(6)"`             //道具打包数量
	Weight   int   `db:"int(11)"`                 //权重
	VipLv    int16 `db:"smallint(6)"`             //购买vip等级条件
	Price    int64 `db:"bigint(20)"`              //购买价格
	OriPrice int64 `db:"bigint(20)"`              //商品原价
	BuyLimit int16 `db:"smallint(6)"`             //购买限制,-1为无限购买，0为售罄，>0为有购买数量限制
	Hidden   int16 `db:"smallint(6)"`             //是否隐藏，0为显示，1为隐藏
}
