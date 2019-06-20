package main

import (
	"image"
	"image/png"
	"os"

	imageProcess "github.com/uehr/kokun/pkg/imageProcess"
	// senryu "github.com/uehr/kokun/pkg/senryu"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	fontSize := 30.0

	imageProcess.AddVerticalLabel(img, 0, 0, "五時過ぎた カモンベイベー USAばらし", "fonts-japanese-mincho.ttf", fontSize)

	f, err := os.Create("hello-go.png")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
