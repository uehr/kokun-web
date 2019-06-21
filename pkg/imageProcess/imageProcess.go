package imageProcess

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// フォント読み込み
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

// 横書きの文字を入れる
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

func AddVerticalBottomAlignLabel(img *image.RGBA, leftX int, label, fontPath string, fontSize float64, fontColor color.Color, marginPx int) error {
	size := GetSize(img)
	labelLength := utf8.RuneCountInString(label)
	labelHeight := labelLength * int(fontSize)
	y := size.Y - labelHeight - marginPx

	err := AddVerticalLabel(img, leftX, y, label, fontPath, fontSize, fontColor)
	if err != nil {
		return err
	}

	return nil
}

func AddVerticalCenterAlignLabel(img *image.RGBA, leftX int, label, fontPath string, fontSize float64, fontColor color.Color, marginPx int) error {
	size := GetSize(img)
	labelLength := utf8.RuneCountInString(label)
	labelHeight := labelLength * int(fontSize)
	padding := size.Y - marginPx*2 - labelHeight
	y := padding / 2

	err := AddVerticalLabel(img, leftX, y, label, fontPath, fontSize, fontColor)
	if err != nil {
		return err
	}

	return nil
}

func AddVerticalTopAlignLabel(img *image.RGBA, leftX int, label, fontPath string, fontSize float64, fontColor color.Color, marginPx int) error {
	y := marginPx

	err := AddVerticalLabel(img, leftX, y, label, fontPath, fontSize, fontColor)
	if err != nil {
		return err
	}

	return nil
}

func GetSize(img *image.RGBA) image.Point {
	return img.Bounds().Size()
}

// 縦書きの文字を入れる
func AddVerticalLabel(img *image.RGBA, leftX, topY int, label, fontPath string, fontSize float64, fontColor color.Color) error {
	charBottomY := topY + int(fontSize)

	for _, char := range label {
		err := AddHorizontalLabel(img, leftX, charBottomY, string(char), fontPath, fontSize, fontColor)

		if err != nil {
			return err
		}

		charBottomY += int(fontSize)
	}

	return nil
}

//画像を保存
func SaveImage(img *image.RGBA, filePath string) error {
	//ファイル作成と後始末
	file, err := os.Create(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	// PNG出力
	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}

// 背景色を追加
func SetBackgroundColor(img *image.RGBA, color image.Image) {
	draw.Draw(img, img.Bounds(), color, image.ZP, draw.Src)
}

// 新規画像作成
func NewImage(width int, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// 外枠を作成
// func AddBorder(img *image.RGBA, color image.Image, thicknessPx int) {
// 	size := GetSize(img)
// }

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(img *image.RGBA, x1, y1, x2, y2, thicknessPx int, col color.Color) {
	for i := 0; i < thicknessPx; i++ {
		HLine(img, x1+i, y1+i, x2-i, col)
		HLine(img, x1+i, y2-i, x2-i, col)
		VLine(img, x1+i, y1+i, y2-i, col)
		VLine(img, x2-i, y1+i, y2-i, col)
	}
}
