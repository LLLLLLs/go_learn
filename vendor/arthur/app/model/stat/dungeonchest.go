/*
Created on 2018-09-07 09:59:26
author: Auto Generate
*/
package stat

type DungeonChest struct {
	No        int16    `model:"pk"` //宝箱编号,1.青铜;2.白银;3.黄金
	Cost      int      //积分消耗
	Key       int16    //所需钥匙编号
	AwardList []string //奖励列表
}
