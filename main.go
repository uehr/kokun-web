package main

import (
	"flag"
	"fmt"

	imageProcess "github.com/uehr/kokun/pkg/imageProcess"
	senryu "github.com/uehr/kokun/pkg/senryu"
)

func main() {
	flag.Parse()
	sen := senryu.Senryu{}

	sen.FirstSentence = flag.Arg(0)
	sen.SecondSentence = flag.Arg(1)
	sen.ThirdSentence = flag.Arg(2)
	sen.AuthorName = flag.Arg(3)

	option := senryu.SenryuImageOption{}
	senryuImage, err := senryu.CreateImage(&sen, &option)

	imageProcess.PasteImage(senryuImage, 100, 100, "hoge.jpg")

	if err != nil {
		fmt.Println("error:file\n", err)
		return
	}

	imageProcess.SaveImage(senryuImage, "senryu.png")
}
