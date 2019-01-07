/*
Created on 2018-10-22 14:57:33
author: Auto Generate
*/
package stat

type VisitMainStory struct {
	No              int      `model:"pk"` //剧情编号
	Country         int      //触发国家
	Award           [][2]int //奖励组
	ChapterRequired int      //关卡要求
}
