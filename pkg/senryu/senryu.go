package senryu

import (
	"image"
	"image/color"

	"github.com/uehr/kokun/pkg/imageProcess"
)

const DefaultSenryuHeight = 1000
const DefaultSenryuWidth = 400
const DefaultFirstSentenceTopY = 0
const DefaultFirstSentenceLeftX = 0
const DefaultSecondSentenceTopY = 0
const DefaultSecondSentenceLeftX = 0
const DefaultThirdSentenceTopY = 0
const DefaultThirdSentenceLeftX = 0
const DefaultFontSize = 60
const DefaultFontPath = "fonts-japanese-gothic.ttf"

var DefaultBackgroundColor = image.White
var DefaultFontColor = color.Black
var DefaultBorderColor = color.RGBA{0, 255, 0, 255}

type Senryu struct {
	FirstSentence  string
	SecondSentence string
	ThirdSentence  string
	AuthorName     string
}

type SenryuImageOption struct {
	SenryuHeight        int
	SenryuWidth         int
	FirstSentenceTopY   int
	FirstSentenceLeftX  int
	SecondSentenceTopY  int
	SecondSentenceLeftX int
	ThirdSentenceTopY   int
	ThirdSentenceLeftX  int
	FontSize            float64
	FontPath            string
	FontColor           color.Color
	BackgroundColor     image.Image
}

func CompleteSenryuImageOption(option *SenryuImageOption) {
	// 変数に値が代入されていないとデフォルト値で補完
	if option.SenryuHeight == 0 {
		option.SenryuHeight = DefaultSenryuHeight
	}

	if option.SenryuWidth == 0 {
		option.SenryuWidth = DefaultSenryuWidth
	}

	if option.FirstSentenceLeftX == 0 {
		option.FirstSentenceLeftX = DefaultFirstSentenceLeftX
	}

	if option.FirstSentenceTopY == 0 {
		option.FirstSentenceTopY = DefaultFirstSentenceTopY
	}

	if option.SecondSentenceLeftX == 0 {
		option.SecondSentenceLeftX = DefaultSecondSentenceLeftX
	}

	if option.SecondSentenceTopY == 0 {
		option.SecondSentenceTopY = DefaultSecondSentenceTopY
	}
	if option.ThirdSentenceLeftX == 0 {
		option.ThirdSentenceLeftX = DefaultThirdSentenceLeftX
	}

	if option.ThirdSentenceTopY == 0 {
		option.ThirdSentenceTopY = DefaultThirdSentenceTopY
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

	if option.BackgroundColor == nil {
		option.BackgroundColor = DefaultBackgroundColor
	}

}

func CreateImage(s *Senryu, option *SenryuImageOption) (*image.RGBA, error) {
	CompleteSenryuImageOption(option)

	img := imageProcess.NewImage(option.SenryuWidth, option.SenryuHeight)

	imageProcess.SetBackgroundColor(img, option.BackgroundColor)

	err := imageProcess.AddVerticalLabel(img, option.FirstSentenceLeftX, option.FirstSentenceTopY, s.FirstSentence, option.FontPath, option.FontSize, option.FontColor)

	if err != nil {
		return nil, err
	}

	err = imageProcess.AddVerticalLabel(img, option.SecondSentenceLeftX, option.SecondSentenceTopY, s.SecondSentence, option.FontPath, option.FontSize, option.FontColor)

	if err != nil {
		return nil, err
	}

	err = imageProcess.AddVerticalLabel(img, option.ThirdSentenceLeftX, option.ThirdSentenceTopY, s.ThirdSentence, option.FontPath, option.FontSize, option.FontColor)

	if err != nil {
		return nil, err
	}

	return img, nil
}
