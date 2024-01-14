package utils

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"image/color"
)

var colors = map[string]color.Color{
	"pink":   color.CMYK{C: 1, M: 167, Y: 111, K: 83},
	"blue":   color.CMYK{C: 182, M: 79, K: 178},
	"brown":  color.CMYK{M: 63, Y: 111, K: 170},
	"orange": color.CMYK{M: 137, Y: 220, K: 12},
}

func GetColor(userColor string) color.Color {
	c, ok := colors[userColor]
	if !ok {
		c = color.Transparent
	}
	return c
}

func QrGenerate(userColor, url string, path string) {
	c := GetColor(userColor)

	err := qrcode.WriteColorFile(
		url,
		qrcode.Medium, 128,
		color.Transparent,
		c,
		path+"/qr.png",
	)

	if err != nil {
		fmt.Println("error", err)
	}
}
