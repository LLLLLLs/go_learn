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
	simpleClick = make([]gocv.Mat, 0)
	upgradeTmpl gocv.Mat
	refreshTmpl gocv.Mat
)

var (
	skillPriority = []string{}
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
	upgradeTmpl = gocv.IMRead(ResourcePath.Join("refresh.png").String(), gocv.IMReadColor)
}
