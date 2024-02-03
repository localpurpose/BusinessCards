package utils

import (
	"github.com/gofiber/fiber/v2/log"
	qrcode "github.com/skip2/go-qrcode"
	"image/color"
)

var colors = map[string]color.Color{
	"pink":   color.CMYK{1, 167, 111, 83},
	"blue":   color.CMYK{182, 79, 0, 178},
	"brown":  color.CMYK{0, 63, 111, 170},
	"orange": color.CMYK{0, 137, 220, 12},
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
		log.Info(err)
	}
}
