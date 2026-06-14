package generator

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

// カート本体の緑色。RGBA() は 16bit 値を返すため RGB(146,208,80) はこの値になる。
const (
	cartGreenR = 37522
	cartGreenG = 53456
	cartGreenB = 20560
)

func (g *Generator) GenerateImage(text string, colorStr string) (*gif.GIF, error) {
	// 色未指定なら元の緑のまま塗り替えない。
	var newColor *color.RGBA
	if colorStr != "" {
		rgbs := strings.Split(colorStr, ",")
		if len(rgbs) != 3 {
			return nil, fmt.Errorf("invalid color %q: want \"R,G,B\"", colorStr)
		}
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
		newColor = &color.RGBA{R: uint8(r), G: uint8(gv), B: uint8(b), A: 255}
	}

	for i, frame := range g.Base.Image {
		// パレット GIF では Set が最近傍色へスナップするため、緑ピクセルを塗り直すのではなく
		// パレットの緑エントリ自体を要求色へ差し替えて、要求色を正確に反映させる。
		if newColor != nil {
			for idx, c := range frame.Palette {
				if cr, cg, cb, _ := c.RGBA(); cr == cartGreenR && cg == cartGreenG && cb == cartGreenB {
					frame.Palette[idx] = *newColor
				}
			}
		}

		var x, y int
		cs := getStringArrayByText(text)

		// 1
		x, y = getFirstCharacterPosition(i)
		g.addCharactorToImage(frame, x, y, cs[0])
		// 2
		x, y = getSecondCharacterPosition(i)
		g.addCharactorToImage(frame, x, y, cs[1])
		// 3
		x, y = getThirdCharacterPosition(i)
		g.addCharactorToImage(frame, x, y, cs[2])
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
