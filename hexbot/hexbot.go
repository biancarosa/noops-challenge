package main

import (
	"fmt"
	"github.com/biancarosa/noops-challenge/colors"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	//Create image
	width := 600
	height := 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	hex := colors.GetHexColor()
	blockW := width / len(hex.Colors)
	blockH := height / len(hex.Colors)
	rgba, _ := colors.ParseHexColor(hex.Colors[len(hex.Colors)/2].Value)
	draw.Draw(img, img.Bounds(), &image.Uniform{rgba}, image.ZP, draw.Src)

	for i, color := range hex.Colors {
		rgba, _ := colors.ParseHexColor(color.Value)
		x0 := blockW * i
		y0 := blockH * i
		x1 := blockW * (i + 1)
		y1 := blockH * (i + 1)
		fmt.Println(i, " - ", x0, y0, x1, y1)
		mask := image.Rect(x0, y0, x1, y1)
		draw.Draw(img, mask, &image.Uniform{rgba}, image.ZP, draw.Src)
	}

	// Save to out.png
	f, _ := os.OpenFile("images/rectangle.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
