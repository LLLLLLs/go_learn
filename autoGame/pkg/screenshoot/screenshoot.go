package screenshoot

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func Screenshoot() {
	for i := range 2 {
		bounds := screenshot.GetDisplayBounds(i)
		fmt.Printf("屏幕%d分辨率: %v\n", i, bounds)
	}
	// 设置截取区域：从 (x=100, y=100)，宽度 300，高度 200
	rect := image.Rect(40, 150, 420, 780)

	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}

	// 保存到文件
	file, err := os.Create("../../resources/partial_capture.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)

	fmt.Println("截取完成，保存为 partial_capture.png")
}
