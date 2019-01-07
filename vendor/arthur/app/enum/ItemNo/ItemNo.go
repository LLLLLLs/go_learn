package ItemNo

//道具编号
type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

//特殊道具， ValNo = 0
const (
	TianFuDianDai Type = 1  //天赋点袋
	GaiMingKa     Type = 2  //改名卡
	HuoLiDan      Type = 3  // 活力丹
	ArenaBaoZhu   Type = 6  //竞技场宝珠
	XingDongYaoJi Type = 8  //行动药剂
	WeiPaiShu     Type = 9  //委派书
	HuoLiYaoJi    Type = 10 //活力药剂
	ShuaiXinLing  Type = 11 //刷新令
	ZhengShouLing Type = 12 // 征收令
	ChuShiLing    Type = 13 // 出使令
	TiaoZhanShu   Type = 14 // 挑战书
	NanJueSuiPian Type = 15 //男爵碎片
	YaoYueHai     Type = 17 //邀约函

	BronzeKey   Type = 18 // 青铜钥匙
	SilverKey   Type = 19 // 白银钥匙
	GoldenKey   Type = 20 // 黄金钥匙
	ChuZhanLing Type = 21 // 出战令

	// 联姻道具
	SilverRing  Type = 26 // 银戒指
	GoldenRing  Type = 27 // 金戒指
	DiamondRing Type = 28 // 钻石戒指

	ExamBook     Type = 30 // 修道院考校宝典
	StudentMedal Type = 31 // 学员奖章 --> 指定鼓励用

	YuanZhuoZhaoLing Type = 32
)

var DungeonKeys = []Type{BronzeKey, SilverKey, GoldenKey}

//奖励类道具， ValNo = 1
const (
	YinPiao       Type = 101 //银票
	LiangPiao     Type = 102 //粮票
	BingFu        Type = 103 //兵符
	MingWangLiBao Type = 104 //名望礼包
	MingWangLiHe  Type = 105 //名望礼盒
)

const (
	TuShangLeiYao   Type = 301 //土商类要
	ZhongChengShiJi Type = 322 //忠诚试剂
	YingyongShiJi   Type = 323 //英勇试剂
)

const (
	WisdomDiamond   Type = 401 //智慧宝石
	DiligentDiamond Type = 402 //勤勉宝石
	LoyalDiamond    Type = 403 //忠诚宝石
	BraveDiamond    Type = 404 //英勇宝石
)

//升爵道具, ValNo = 5
const (
	BaronDuanJian Type = 501
	BaronXunZhang Type = 502
	BaronGuanMian Type = 503

	ViscountDuanJian Type = 511
	ViscountXunZhang Type = 512
	ViscountGuanMian Type = 513
)

//beauty item

const (
	KuiJia  = 701 //盔甲
	BaoJian = 702 //宝剑
	ZhuBao  = 703 //珠宝
	HuaFu   = 704 //华服
	//711-745 1-35号名媛碎片
	// 名媛碎片起始id
	StartBeautyNo = 710
)

const (
	HeroPiece33 = 933 // 33号英雄碎片
)
