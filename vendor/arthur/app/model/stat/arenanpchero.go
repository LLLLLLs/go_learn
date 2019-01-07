/*
Created on 2018-10-24 11:35:47
author: Auto Generate
*/
package stat

type ArenaNpcHero struct {
	NpcId   string `db:"char(36)"`    //所属npc
	No      int16  `db:"smallint(5)"` //英雄编号
	Attack  int64  `db:"bigint(20)"`  //英雄攻击
	Blood   int64  `db:"bigint(20)"`  //英雄血量
	Peerage int16  `db:"smallint(5)"` //英雄爵位
	Level   int16  `db:"smallint(5)"` //英雄等级
}
