package main

import (
	"fmt"
	"github.com/biancarosa/noops-challenge/colors"
	"github.com/biancarosa/noops-challenge/pictures"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func combine(c1, c2 color.Color) color.Color {
	r, g, b, a := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return color.RGBA{
		uint8((r*2 + r2) / 3 >> 8),
		uint8((g*2 + g2) / 3 >> 8),
		uint8((b*2 + b2) / 3 >> 8),
		uint8((a*2 + a2) / 3 >> 8),
	}
}

// Color strips on a random image
func main() {
	pic := pictures.GetPicture()
	img, _ := pictures.DownloadPicture(pic)

	drawable := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	hex := colors.GetHexColor()
	blockW := drawable.Bounds().Dx() / len(hex.Colors)
	draw.Draw(drawable, drawable.Bounds(), img, image.ZP, draw.Src)

	for i, color := range hex.Colors {
		rgba, _ := colors.ParseHexColor(color.Value)
		x0 := blockW * i
		x1 := blockW * (i + 1)
		y0 := 0
		y1 := drawable.Bounds().Dy()
		fmt.Println(i, " - ", x0, y0, x1, y1)
		for x := x0; x <= x1; x++ {
			for y := y0; y <= y1; y++ {
				drawable.Set(x, y, combine(img.At(x, y), rgba))
			}
		}
	}

	// Save to out.png
	f, _ := os.OpenFile("images/stripped.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, drawable)
}
