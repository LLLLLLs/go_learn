//errors
//created: 2018/8/16
//author: wdj

package errors

var (
	ErrNoSuchPos             = New("invalid_position")       //地点不存在
	ErrNoSuchCity            = New("city_not_exist")         //城堡不存在
	ErrBossPosNotVisited     = New("boss_pos_not_visited")   //boss点未访问
	ErrHeroNotDispatched     = New("hero_not_dispatched")    //英雄未派遣
	ErrMaxCity               = New("max_city")               //达到最大城堡数
	ErrInvalidAward          = New("invalid_award")          //非法奖励
	ErrInsufficientAdventure = New("insufficient_adventure") //探索点不足
	ErrPosNotVisited         = New("position_not_visited")   //未进入地点
	ErrMoreCityRequired      = New("more_city_required")     //解锁城堡不足
)
