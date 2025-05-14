package internal

import (
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
	if upgradeCheck(img) {
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
	clickPoint(maxLoc, tmpl)
	return true
}

func clickPoint(start image.Point, tmpl gocv.Mat) {
	center := image.Point{
		X: start.X + tmpl.Cols()/2 + 40,
		Y: start.Y + tmpl.Rows()/2 + 150,
	}
	Click(center.X, center.Y)
}

func matchResult(img, tmpl gocv.Mat) (maxLoc image.Point, matched bool) {
	result := gocv.NewMat()
	defer result.Close()
	gocv.MatchTemplate(img, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	var maxVal float32
	_, maxVal, _, maxLoc = gocv.MinMaxLoc(result)
	matched = maxVal > 0.95
	return
}

func upgradeCheck(img gocv.Mat) bool {
	if _, matched := matchResult(img, upgradeTmpl); !matched {
		return false
	}
	canRefresh = nil
	for i := range skillTmpl {
		if i > len(skillTmpl)*2/3 && refresh(img) {
			return true
		}
		maxLoc, matched := matchResult(img, skillTmpl[i])
		if matched {
			clickPoint(maxLoc, skillTmpl[i])
		}
	}
	return true
}

var refreshLoc image.Point
var canRefresh *bool

func refresh(img gocv.Mat) bool {
	if canRefresh == nil {
		refreshLoc, *canRefresh = matchResult(img, refreshTmpl)
	}
	if !(*canRefresh) {
		return false
	}
	clickPoint(refreshLoc, refreshTmpl)
	return true
}
