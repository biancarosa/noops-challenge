package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
)

type Hex struct {
	Colors []struct {
		Value       string
		Coordinates *string
	}
}

//Stole from here: https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func main() {
	//Get hexcolor
	response, _ := http.Get("https://api.noopschallenge.com/hexbot")
	var hex Hex
	buf, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(buf, &hex)
	fmt.Printf("%#v\n", hex)

	//Create image
	width := 600
	height := 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for _, color := range hex.Colors {
		rgba, _ := ParseHexColor(color.Value)
		draw.Draw(img, img.Bounds(), &image.Uniform{rgba}, image.ZP, draw.Src)
	}

	// Save to out.png
	f, _ := os.OpenFile("images/rectangle.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
