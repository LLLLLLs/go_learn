/*
Created on 2018-10-19 16:32:36
author: Auto Generate
*/
package stat

type ArenaLevelBuff struct {
	No           int16  `model:"pk"` //
	AttackLower  int16  //攻击最低加成
	AttackHeight int16  //攻击最高加成
	BloodLower   int16  //血量最低加成
	BloodHeight  int16  //血量最高加成
	Remark       string //备注
}
