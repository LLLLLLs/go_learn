package main

import (
	"os"
	"runtime"
)

func configureFyneFont() {
	if os.Getenv("FYNE_FONT") != "" {
		return
	}

	for _, candidate := range fontCandidates() {
		if _, err := os.Stat(candidate); err == nil {
			_ = os.Setenv("FYNE_FONT", candidate)
			return
		}
	}
}

func fontCandidates() []string {
	switch runtime.GOOS {
	case "darwin":
		return []string{
			"/System/Library/Fonts/Supplemental/Arial Unicode.ttf",
			"/System/Library/Fonts/STHeiti Medium.ttc",
		}
	case "windows":
		return []string{
			`C:\Windows\Fonts\msyh.ttc`,
			`C:\Windows\Fonts\simhei.ttf`,
		}
	default:
		return []string{
			"/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
			"/usr/share/fonts/truetype/wqy/wqy-microhei.ttc",
		}
	}
}
