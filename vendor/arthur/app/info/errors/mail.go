/*
Author: Minsi Ruan
Data: 2018/7/19 15:35
*/

package errors

import "arthur/utils/errors"

var (
	NotExistMailError    = errors.New("err_doesnot_exist_mail") //邮件不存在
	UnShieldError        = errors.New("err_unshield")           //邮件未屏蔽
	MailShieldError      = errors.New("err_shield")             //邮件已屏蔽
	AlreadyReceiveError  = errors.New("err_already_receive")    //邮件已领取
	NotReadMailError     = errors.New("err_not_read_mail")      //没有可读邮件
	NotReceivedMailError = errors.New("err_not_received_mail")  //没有可领取的邮件
)
