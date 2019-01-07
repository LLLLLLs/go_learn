package chat

const (
	SUCCESS         = 1   //成功
	PARAM_MISSING   = 100 //缺少必要参数
	PARAM_ILLEGAL   = 101 //非法参数
	IP_ILLEGAL      = 102 //非法IP
	OVERTIME        = 103 //请求超时
	LIMITED_FLOW    = 104 //请求限流
	MSG_EMPTY       = 480 //	消息为空
	DB_CONN_FAILED  = 481 //数据库连接失败
	UNKNOW_ERROR    = 482 //未知错误
	NOT_AUTHORIZED  = 483 //未认证请求
	ADVERTISEMENT   = 484 //广告信息
	BAD_SPEECH      = 485 //不良言论
	FORBIDDEN_SPEAK = 486 //被禁言
	NOT_CHAT_CODE   = 999 //非聊天错误
)
