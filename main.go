package main

import (
	"bytes"
	"encoding/base64"
	"image/gif"
	"syscall/js"

	"example.com/countdown-gif/generator"
)

func generate(this js.Value, args []js.Value) interface{} {
	// 入力値取得
	word := args[0].String()
	color := args[1].String()

	generator, err := generator.NewGenerator()
	if err != nil {
		return js.ValueOf(err.Error())
	}

	img, err := generator.GenerateImage(color, word)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	// GIF生成
	var buf bytes.Buffer
	if err := gif.EncodeAll(&buf, img); err != nil {
		return js.ValueOf(err.Error())
	}

	// Base64エンコードして返す
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return js.ValueOf(encoded)
}

func main() {
	js.Global().Set("generate", js.FuncOf(generate))
	select {}
}
