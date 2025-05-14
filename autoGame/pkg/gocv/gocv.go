package gocv

import (
	"autogame/pkg/util"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"log"
)

const ResourcePath = util.Path("../../resources")

func CV() {
	tmplPath := ResourcePath.Join("fight.png").String() // 模板图片
	tmpl := gocv.IMRead(tmplPath, gocv.IMReadColor)
	if tmpl.Empty() {
		panic("无法读取模板图像")
	}

	img := util.Screenshoot()

	result := gocv.NewMat()
	gocv.MatchTemplate(img, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	if maxVal > 0.9 {
		rect := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+tmpl.Cols(), maxLoc.Y+tmpl.Rows())
		// 选取该区域
		cropped := img.Region(rect)
		defer cropped.Close()
		// 保存裁剪区域为新文件
		outputPath := ResourcePath.Join("cropped.png").String()
		ok := gocv.IMWrite(outputPath, cropped)
		if !ok {
			log.Fatalf("无法写入文件: %s", outputPath)
		}
		fmt.Println("裁剪区域保存成功:", outputPath)
		center := image.Point{
			X: maxLoc.X + tmpl.Cols()/2 + 40,
			Y: maxLoc.Y + tmpl.Rows()/2 + 150,
		}
		fmt.Printf("找到目标，点击中心: %+v\n", center)
		util.Click(center.X, center.Y)
	}

	tmpl.Close()
	img.Close()
	result.Close()
}
