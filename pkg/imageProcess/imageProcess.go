// package imageProcess
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func fontload(fname string) []byte {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		fmt.Println("error:file\n", err)
		return nil
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:fileread\n", err)
		return nil
	}

	return bytes
}

func AddHorizontalLabel(img *image.RGBA, leftX, bottomY int, label, fontPath string, fontSize float64, fontColor color.Color) error {
	ft, err := truetype.Parse(fontload(fontPath))
	if err != nil {
		return err
	}

	opt := truetype.Options{Size: fontSize}
	face := truetype.NewFace(ft, &opt)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fontColor),
		Face: face,
	}

	d.Dot.X = fixed.I(int(leftX))
	d.Dot.Y = fixed.I(int(bottomY))

	d.DrawString(label)

	return nil
}

func AddVerticalLabel() error {
	return nil
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	err := AddHorizontalLabel(img, 0, 50, "テストだよ", "fonts-japanese-gothic.ttf", 30.0, color.RGBA{0, 255, 0, 255})

	if err != nil {
		fmt.Println("error:file\n", err)
		return
	}

	//ファイル作成と後始末
	file, err := os.Create("test.png")
	defer file.Close()
	if err != nil {
		fmt.Println("error:file\n", err)
		return
	}

	// PNG出力
	if err := png.Encode(file, img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
