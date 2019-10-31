// Time        : 2019/10/21
// Description :

package leaderboards

type DefaultEvent struct {
	Type     uint16      `rt:"type"`
	Identify string      `rt:"identity"`
	Value    int64       `rt:"value"`
	Extend   interface{} `rt:"extend"`
}

// identify::group::group
// group的设定:
// 	- 主要用来做队伍、联盟等成员的排行 team1::role2 alliance3::role4
//	- 可以用英雄或名媛等次级目标来作为排行榜目标 role::1 role::2
//	- 成就或/任务可能要记录玩家不同等级的累计数据 role::1 role::2 role::3 这样就可以分别累计玩家在指定状态下的数据聚合
type CustomerEvent struct {
	Type     uint16      `rt:"type"`
	Identify string      `rt:"identity"`
	Group    string      `rt:"group"`
	Group2   int64       `rt:"group"`
	Value    int64       `rt:"value"`
	Extend   interface{} `rt:"extend"`
}

type AnotherEvent struct {
	A uint16   `rt:"type"`
	B string   `rt:"identity"`
	C string   `rt:"group"`
	D int64    `rt:"group"`
	E int64    `rt:"value"`
	F struct{} `rt:"extend"`
}

// mongo
type RunningTotal struct {
	Id       string
	Type     uint16
	BelongTo string
	Value    int64
	Extend   interface{}
}

// running total rule
type RunningTotalRule struct {
	Type       uint16
	GatherType GatherType
	Reset      ResetType // ???
}

type GatherType uint

const (
	Sum GatherType = 0 // default
	Max GatherType = 1
	Min GatherType = 2
	// ...
)

type SortType uint

const (
	DESC SortType = 0 // default
	ASC  SortType = 1
)

type ResetType uint

const (
	Never   ResetType = 0 // default
	Daily   ResetType = 1
	Weekly  ResetType = 2
	Monthly ResetType = 3
)

// redis
type LeaderBoard struct {
}

type LeaderBoardRule struct {
	Type uint16
	Sort SortType
}
