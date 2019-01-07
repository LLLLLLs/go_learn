package ItemType

//道具类型
type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Consume     Type = 0 //特殊道具，每个道具有单独的使用接口
	Award       Type = 1 //使用后会获得奖励的道具
	HeroAttr    Type = 3 //英雄属性道具
	HeroTalent  Type = 4 //英雄天赋道具
	HeroPeerage Type = 5 //英雄升爵道具
	Fashion     Type = 6 //时装
	Beauty      Type = 7
	BeautyPiece Type = 8
	HeroPiece   Type = 9
)
