//@author: lls
//@time: 2020/11/12
//@desc:

package subject

import "fmt"

// 设想有一个喜剧演出团，客户会指定几出剧目，剧团根据观众人数及剧目类型来向客户收费。
// 该团目前有两种戏剧：悲剧(tragedy)和戏剧(comedy)。给客户发出账单时，
// 剧团还会根据到场观众的数量给出"观众量积分"优惠，下次客户再请剧团表演时可以使用积分获得折扣。
// ************ 功能扩展 ***********
// 1.输出类型：plain text、html等
// 2.剧目类型：新增

type Play struct {
	Name string // 剧目名称
	Type string // 剧目类型 "tragedy" "comedy"
}

type Plays map[string]Play // 剧目列表 id->play

type Invoice struct {
	Customer     string        // 客户名称
	Performances []Performance // 演出数据
}

type Performance struct {
	PlayId   string // 剧目id 关联plays
	Audience int    // 观众数量
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func USD(cent int) string {
	d := float64(cent) / 100
	return fmt.Sprintf("$%.02f", d)
}
