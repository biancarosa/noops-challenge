package colors

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
)

type Hex struct {
	Colors []struct {
		Value       string
		Coordinates *string
	}
}

func GetHexColor() Hex {
	//Get hexcolor
	resp, _ := http.Get("https://api.noopschallenge.com/hexbot?count=20")
	var hex Hex
	buf, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(buf, &hex)
	fmt.Printf("%#v\n", hex)
	return hex
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
