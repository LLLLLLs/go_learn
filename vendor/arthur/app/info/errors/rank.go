package errors

var (
	ErrOutOfRank         = New("out_of_rank")         //超出排行
	ErrNoRank            = New("no_rank")             //找不到排行
	ErrCannotWorshipRank = New("cannot_worship_rank") //无法膜拜
	ErrRankNoRole        = New("rank_no_role")        //无人上榜
	ErrHasWorship        = New("has_worship")         //已膜拜
)
