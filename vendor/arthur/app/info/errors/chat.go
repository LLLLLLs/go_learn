/*
Created on 2018/8/30 13:42

author: ChenJinLong

Content:
*/
package errors

var (
	ErrChatReportReasonWordsExceed   = New("report_reason_words_exceed")   //举报理由字数超过上限
	ErrChatMessageIdNotExist         = New("message_id_not_exist")         //没有此消息
	ErrChatCantReportYourself        = New("cant_report_yourself")         //不能举报自己
	ErrChatHaveReportedToday         = New("have_reported_today")          //今天已经举报
	ErrChatCantBlockYourself         = New("cant_block_yourself")          //不能屏蔽自己
	ErrChatBlockPlayersBeyondLimit   = New("block_players_beyond_limit")   //屏蔽人数达到上限
	ErrChatPlayerHasBeenBlocked      = New("player_has_been_blocked")      //玩家被屏蔽
	ErrChatAllowedTalkLvInsufficient = New("allowed_talk_lv_insufficient") //等级未到，不能发言
	ErrChatReportTimeNotArrived      = New("report_time_not_arrived")      //举报间隔未到
	ErrChatShareTimeNotArrived       = New("share_time_not_arrived")       //分享间隔未到
	ErrChatNullText                  = New("null_text")                    //空文本
	ErrChatWordsLimit                = New("words_limit")                  //字数超过上限
	ErrChatFirstWordBlank            = New("first_word_blank")             //首字母为空
	ErrChatAwardEmpty                = New("chat_award_empty")             //奖励已领完
	ErrChatGainAwardBySelf           = New("can_not_gain_award_by_self")   //不能领取自己的奖励
	ErrParamMissing                  = New("param_missing")                //缺少必要参数
	ErrIpIllegal                     = New("ip_illegal")                   //非法IP
	ErrOvertime                      = New("overtime")                     //请求超时
	ErrLimitedFlow                   = New("limited_flow")                 //请求限流
	ErrMsgEmpty                      = New("msg_empty")                    //	消息为空
	ErrDbConnFailed                  = New("db_conn_failed")               //数据库连接失败
	ErrUnknowError                   = New("unknow_error")                 //未知错误
	ErrNotAuthorized                 = New("not_authorized")               //未认证请求
	ErrAdvertisement                 = New("advertisement")                //广告信息
	ErrBadSpeech                     = New("bad_speech")                   //不良言论
	ErrForbiddenSpeak                = New("forbidden_speak")              //被禁言New(")
)
