package tesseract

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"image/png"
)

func Tesseract() {
	// 加载已有截图（例如 screen.png）
	srcImagePath := "../../resources/defeat.png"
	img, err := imaging.Open(srcImagePath)
	if err != nil {
		panic(err)
	}

	// 2. 裁剪区域：左上(100,100) 到 右下(400,300)
	//cropRect := image.Rect(100, 100, 400, 300)
	//cropped := imaging.Crop(img, cropRect)

	// 3. 将裁剪后的图像编码为 PNG（字节流）
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}
	imgBytes := buf.Bytes()

	// 4. 初始化 OCR 客户端
	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("chi_sim") // 识别中文
	err = client.SetImageFromBytes(imgBytes)
	if err != nil {
		panic(err)
	}

	// 5. 获取 OCR 结果
	text, err := client.Text()
	if err != nil {
		panic(err)
	}

	fmt.Println("识别结果：")
	fmt.Println("---------------------")
	fmt.Println(text)
	fmt.Println("---------------------")
}
