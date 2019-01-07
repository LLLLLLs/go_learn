/*
Author      : lls
Time        : 2018/08/21
Description : fashion errors
*/

package errors

var (
	ErrNoFashionConf   = New("no_fashion")        // 无时装配置
	ErrCanNotEdit      = New("can_not_edit")      // 时装不可编辑宣言
	ErrSentenceTooLong = New("sentence_too_long") // 宣言过长
	ErrSentenceIllegal = New("sentence_illegal")  // 宣言不合法
	ErrNoFashion       = New("fashion_unlock")    // 时装未解锁
	ErrFashionOutDated = New("fashion_out_date")  // 时装已过期
)
