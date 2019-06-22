package senryu

import (
	"image"
	"image/color"

	"github.com/uehr/kokun/pkg/imageProcess"
)

const DefaultSenryuHeight = 1000
const DefaultSenryuWidth = 550

const DefaultFirstSentenceLeftX = 390
const DefaultSecondSentenceLeftX = 270
const DefaultThirdSentenceLeftX = 160
const DefaultAuthorNameLeftX = 60

const DefaultFontSize = 90
const DefaultAuthorNameFontSize = 50
const DefaultFontPath = "衡山毛筆フォント.ttf"

const DefaultMarginPx = 70
const DefaultThickBorderPx = 50
const DefaultThinBorderPx = 10

var DefaultBackgroundColor = image.White
var DefaultFontColor = color.Black
var DefaultThickBorderColor = color.RGBA{0, 128, 79, 255}
var DefaultThinBorderColor = color.RGBA{255, 215, 0, 255}

type Senryu struct {
	FirstSentence  string
	SecondSentence string
	ThirdSentence  string
	AuthorName     string
}

type SenryuImageOption struct {
	SenryuHeight        int
	SenryuWidth         int
	FirstSentenceLeftX  int
	SecondSentenceLeftX int
	ThirdSentenceLeftX  int
	AuthorNameLeftX     int
	FontSize            float64
	AuthorNameFontSize  float64
	FontPath            string
	FontColor           color.Color
	BackgroundColor     image.Image
	MarginPx            int
	ThickBorderColor    color.Color
	ThinBorderColor     color.Color
	ThickBorderPx       int
	ThinBorderPx        int
}

// オプションを補完
func CompleteSenryuImageOption(option *SenryuImageOption) {
	// 変数に値が代入されていないとデフォルト値で補完
	if option.SenryuHeight == 0 {
		option.SenryuHeight = DefaultSenryuHeight
	}
	if option.SenryuWidth == 0 {
		option.SenryuWidth = DefaultSenryuWidth
	}
	if option.AuthorNameLeftX == 0 {
		option.AuthorNameLeftX = DefaultAuthorNameLeftX
	}
	if option.FirstSentenceLeftX == 0 {
		option.FirstSentenceLeftX = DefaultFirstSentenceLeftX
	}
	if option.SecondSentenceLeftX == 0 {
		option.SecondSentenceLeftX = DefaultSecondSentenceLeftX
	}
	if option.ThirdSentenceLeftX == 0 {
		option.ThirdSentenceLeftX = DefaultThirdSentenceLeftX
	}
	if option.FontColor == nil {
		option.FontColor = DefaultFontColor
	}
	if option.FontPath == "" {
		option.FontPath = DefaultFontPath
	}
	if option.FontSize == 0 {
		option.FontSize = DefaultFontSize
	}
	if option.AuthorNameFontSize == 0 {
		option.AuthorNameFontSize = DefaultAuthorNameFontSize
	}
	if option.BackgroundColor == nil {
		option.BackgroundColor = DefaultBackgroundColor
	}
	if option.ThickBorderColor == nil {
		option.ThickBorderColor = DefaultThickBorderColor
	}
	if option.ThinBorderColor == nil {
		option.ThinBorderColor = DefaultThinBorderColor
	}
	if option.ThickBorderPx == 0 {
		option.ThickBorderPx = DefaultThickBorderPx
	}
	if option.ThinBorderPx == 0 {
		option.ThinBorderPx = DefaultThinBorderPx
	}
	if option.MarginPx == 0 {
		option.MarginPx = DefaultMarginPx
	}

}

func CreateImage(s *Senryu, option *SenryuImageOption) (*image.RGBA, error) {
	// option.BackgroundColor = image.RGBA{255, 242, 179, 255}
	CompleteSenryuImageOption(option)

	img := imageProcess.NewImage(option.SenryuWidth, option.SenryuHeight)

	imageProcess.SetBackgroundColor(img, option.BackgroundColor)

	// err := imageProcess.AddVerticalLabel(img, option.FirstSentenceLeftX, option.FirstSentenceTopY, s.FirstSentence, option.FontPath, option.FontSize, option.FontColor)
	err := imageProcess.AddVerticalTopAlignLabel(img, option.FirstSentenceLeftX, s.FirstSentence, option.FontPath, option.FontSize, option.FontColor, option.MarginPx)
	err = imageProcess.AddVerticalCenterAlignLabel(img, option.SecondSentenceLeftX, s.SecondSentence, option.FontPath, option.FontSize, option.FontColor, option.MarginPx)
	err = imageProcess.AddVerticalBottomAlignLabel(img, option.ThirdSentenceLeftX, s.ThirdSentence, option.FontPath, option.FontSize, option.FontColor, option.MarginPx)
	err = imageProcess.AddVerticalBottomAlignLabel(img, option.AuthorNameLeftX, s.AuthorName, option.FontPath, option.AuthorNameFontSize, option.FontColor, option.MarginPx)

	imageProcess.Rect(img, 0, 0, option.SenryuWidth, option.SenryuHeight, option.ThickBorderPx, option.ThickBorderColor)
	borderThicknessDiff := option.ThickBorderPx - option.ThinBorderPx
	imageProcess.Rect(img, borderThicknessDiff, borderThicknessDiff, option.SenryuWidth-borderThicknessDiff, option.SenryuHeight-borderThicknessDiff, option.ThinBorderPx, option.ThinBorderColor)

	if err != nil {
		return nil, err
	}

	return img, nil
}
