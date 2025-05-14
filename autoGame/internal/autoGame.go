package internal

import (
	"autogame/pkg/util"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"time"
)

func AutoGame() {
	Init()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			tick()
		}
	}
}

func tick() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	img := Screenshoot()
	defer img.Close()

	if clickCheck(img) {
		return
	}
}

func clickCheck(img gocv.Mat) bool {
	for _, tmpl := range simpleClick {
		if clickCheckOnce(img, tmpl) {
			return true
		}
	}
	return false
}

func clickCheckOnce(img, tmpl gocv.Mat) bool {
	maxLoc, matched := matchResult(img, tmpl)
	if !matched {
		return false
	}
	center := image.Point{
		X: maxLoc.X + tmpl.Cols()/2 + 40,
		Y: maxLoc.Y + tmpl.Rows()/2 + 150,
	}
	fmt.Printf("找到目标，点击中心: %+v\n", center)
	util.Click(center.X, center.Y)
	return true
}

func matchResult(img, tmpl gocv.Mat) (maxLoc image.Point, matched bool) {
	result := gocv.NewMat()
	defer result.Close()
	gocv.MatchTemplate(img, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	var maxVal float32
	_, maxVal, _, maxLoc = gocv.MinMaxLoc(result)
	matched = maxVal > 0.9
	return
}

func checkUpgrade(img gocv.Mat) bool {
	if _, matched := matchResult(img, upgradeTmpl); !matched {
		return false
	}
}
