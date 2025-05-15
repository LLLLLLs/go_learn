package internal

import (
	"autogame/pkg/util"
	"bytes"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
	"image"
	"image/png"
	"os/exec"
)

func Screenshoot() gocv.Mat {
	// 设置截取区域：从 (x=100, y=100)，宽度 300，高度 200
	rect := image.Rect(40, 150, 420, 780)

	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	return RGBAImageToMat(img)
}

func RGBAImageToMat(img *image.RGBA) gocv.Mat {
	// 将 image.RGBA 编码为 PNG 字节流
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}

	// 解码为 OpenCV Mat
	mat, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadUnchanged)
	if err != nil {
		panic(err)
	}
	return mat
}

func Click(x, y int) {
	clickRobotGO(x, y)
}

func quartzClick(x, y int) {
	QuartzClick(x, y)
}

func clickApple(x, y int) {
	script := fmt.Sprintf(`tell application "System Events" to click at {%d, %d}`, x, y)
	cmd := exec.Command("osascript", "-e", script)
	err := cmd.Run()
	util.MustOK(err)
}

func clickRobotGO(x, y int) {
	resetX, resetY := robotgo.Location()
	robotgo.MoveClick(x, y)
	robotgo.MoveClick(x, y)
	robotgo.Move(resetX, resetY)
}
