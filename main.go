package main

import (
	"fmt"

	imageProcess "github.com/uehr/kokun/pkg/imageProcess"
	senryu "github.com/uehr/kokun/pkg/senryu"
)

func main() {
	sen := senryu.Senryu{}

	sen.FirstSentence = "first"
	sen.SecondSentence = "second"
	sen.ThirdSentence = "third"

	option := senryu.SenryuImageOption{}
	senryuImage, err := senryu.CreateImage(&sen, &option)

	// err := imageProcess.AddVerticalLabel(img, 50, 50, "テストだーよ", "fonts-japanese-gothic.ttf", 30.0, color.RGBA{255, 0, 0, 255})

	if err != nil {
		fmt.Println("error:file\n", err)
		return
	}

	imageProcess.SaveImage(senryuImage, "senryu.png")
}
