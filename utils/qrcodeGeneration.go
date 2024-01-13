package utils

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"image/color"
)

//const (
//	pink = 1, 167, 111, 83]
//)

//pink = 1, 167, 111, 83
//blue = 182, 79, 0, 178
//brown = 0, 63, 111, 170
//orange = 0, 137, 220, 12

func qrGenerate(userColor string, url string) {
	err := qrcode.WriteColorFile(
		"https://test.amasaetre.ru",
		qrcode.Medium, 128,
		color.Transparent,
		color.CMYK{0, 54, 86, 5},
		"qr.png")
	if err != nil {
		fmt.Println("error", err)
	}

}
