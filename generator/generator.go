package generator

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/gif"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Generator struct {
	Base *gif.GIF
}

func NewGenerator() (*Generator, error) {
	cartBytes, err := base64.StdEncoding.DecodeString(cart)
	if err != nil {
		return nil, err
	}

	image, err := gif.DecodeAll(bytes.NewReader(cartBytes))

	if err != nil {
		return nil, err
	}

	generator := &Generator{
		Base: image,
	}

	return generator, nil
}

func (g *Generator) GenerateImage(text string, colorStr string) (*gif.GIF, error) {
	var rgbStr string

	if colorStr == "" {
		rgbStr = "37522,53456,20560"
	} else {
		rgbStr = colorStr
	}
	rgbs := strings.Split(rgbStr, ",")
	r, err := strconv.Atoi(rgbs[0])
	if err != nil {
		return nil, err
	}
	gv, err := strconv.Atoi(rgbs[1])
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(rgbs[2])
	if err != nil {
		return nil, err
	}

	for i, image := range g.Base.Image {
		bounds := image.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				colorR, colorG, colorB, _ := image.At(x, y).RGBA()

				if colorR == 37522 && colorG == 53456 && colorB == 20560 {
					image.Set(x, y, color.RGBA{
						R: uint8(r),
						G: uint8(gv),
						B: uint8(b),
						A: 1,
					})
				}
			}
		}

		var x, y int
		cs := getStringArrayByText(text)

		// 1
		x, y = getFirstCharacterPosition(i)
		g.addCharactorToImage(image, x, y, cs[0])
		// 2
		x, y = getSecondCharacterPosition(i)
		g.addCharactorToImage(image, x, y, cs[1])
		// 3
		x, y = getThirdCharacterPosition(i)
		g.addCharactorToImage(image, x, y, cs[2])
	}

	return g.Base, nil
}

func (g *Generator) addCharactorToImage(img *image.Paletted, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	ipaFont, _ := getFontFace()
	face := truetype.NewFace(ipaFont, &truetype.Options{Size: 52.0})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}

	d.DrawString(label)
}

func getStringArrayByText(text string) []string {
	arr := []string{}
	textArr := []rune(text)

	var c string

	for i := 0; i < 3; i++ {
		if len(textArr) > i {
			c = string(textArr[i])
		} else {
			c = ""
		}
		arr = append(arr, c)
	}

	return arr
}

func getFirstCharacterPosition(i int) (x, y int) {
	var table = []struct {
		x int
		y int
	}{
		{0, 0},
		{329, 117},
		{253, 117},
		{204, 117},
		{147, 117},
		{90, 117},
		{23, 117},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	return table[i].x, table[i].y
}
func getSecondCharacterPosition(i int) (x, y int) {
	var table = []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 0},
		{358, 117},
		{310, 117},
		{253, 117},
		{186, 117},
		{129, 117},
		{71, 117},
		{0, 0},
		{0, 0},
	}

	return table[i].x, table[i].y
}
func getThirdCharacterPosition(i int) (x, y int) {
	var table = []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{349, 117},
		{282, 117},
		{225, 117},
		{167, 117},
		{76, 117},
		{0, 0},
	}

	return table[i].x, table[i].y
}

func getFontFace() (*truetype.Font, error) {
	// Base64デコード
	fontBytes, err := base64.StdEncoding.DecodeString(fontBase64)
	if err != nil {
		return nil, err
	}
	// truetype.Parse
	ttf, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	return ttf, nil
}
