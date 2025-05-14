package internal

import (
	"autogame/pkg/util"
	"gocv.io/x/gocv"
	"os"
	"path/filepath"
	"strings"
)

const ResourcePath = util.Path("./resources")

var (
	simpleClick []gocv.Mat
	upgradeTmpl gocv.Mat
	refreshTmpl gocv.Mat
)

var (
	skillPriority = []string{
		"月神_大招", "娜娜_大招", "娜娜_天赋", "娜娜_审判之魂", "月神_超载结界", "月神_空间坍缩", "娜娜_星灵速频", "月神_引力空间", "月神_黑洞效应", "神女_光之阵", "神女_御之光", "娜娜_连发星光", "娜娜_星束爆发",
		"娜娜_辉星续能", "娜娜_源值启动", "娜娜_源妙不可言", "塞拉斯_高频鼓舞", "塞拉斯_强化鼓舞", "塞拉斯_恶念轰击", "妮妮_大招", "妮妮_神秘增幅", "神女_迷之光", "神女_凝之阵",
		"娜娜_源源不绝", "娜娜_魔源力场", "妮妮_受宠若惊", "妮妮_声名远扬", "月神_星月连击", "神女_牌中幻兽", "神女_盾之牌", "娜娜_恐惧之落", "神女_泛之光", "妮妮_渐进节拍", "妮妮_穿云天籁",
		"神女_火之光", "神女_破碎之光", "妮妮_近景返场", "塞拉斯_恶念爆炸", "妮妮_震撼心灵", "塞拉斯_恶念汹涌", "神女_幻之轮转", "神女_幻之救赎",
		"塞拉斯_侵蚀异变", "塞拉斯_深度侵蚀", "神女_叠之牌", "神女_力之牌", "神女_幻之强音",
		"月神_虫洞", "月神_虫洞强化", "月神_虫洞电", "月神_虫洞风", "月神_虫洞雷",
	}
	skillTmpl []gocv.Mat
)

func Init() {
	err := filepath.Walk(ResourcePath.Join("click").AbsPath(), func(path string, info os.FileInfo, err error) error {
		util.MustOK(err)
		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}
		simpleClick = append(simpleClick, gocv.IMRead(path, gocv.IMReadColor))
		return nil
	})
	util.MustOK(err)
	upgradeTmpl = gocv.IMRead(ResourcePath.Join("upgrade.png").String(), gocv.IMReadColor)
	refreshTmpl = gocv.IMRead(ResourcePath.Join("refresh.png").String(), gocv.IMReadColor)
	for i := range skillPriority {
		skillTmpl = append(skillTmpl, gocv.IMRead(ResourcePath.Join(skillPriority[i]+".png").AbsPath(), gocv.IMReadColor))
	}
}
