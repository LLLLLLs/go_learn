/*
Created on 2018/8/20 15:55

author: ChenJinLong

Content:
*/
package errors

var (
	ErrStoreNoEnoughVIPLv = New("store_no_enough_vip_lv") //VIP等级不足以购买商品
	ErrStoreBuyLimitOut   = New("store_buy_limit_out")    //商品已售罄
)
