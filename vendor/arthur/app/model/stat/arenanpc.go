/*
Created on 2018-10-24 11:35:43
author: Auto Generate
*/
package stat

type ArenaNpc struct {
	Id       string `model:"pk"` //NPC id
	Name     string //NPC名字
	Alliance string //NPC联盟
	Fashion  int16  //NPC时装
	Scene    int16  //应用场景1=4v4 2=1v1
}
