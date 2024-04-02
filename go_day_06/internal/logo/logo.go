package logo

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	width    = 300
	height   = 300
	FileName = "amazing_logo.png"
	Template = "./static/logo/alpaca.txt"
)

func GenerateLogo() *image.RGBA {
	file, err := os.Open(Template)
	if err != nil {
		log.Fatalln("open template", err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var j int

	blue := color.RGBA{0, 248, 255, 0xff}
	orange := color.RGBA{255, 150, 113, 0xff}
	pink := color.RGBA{214, 93, 177, 0xff}
	purple := color.RGBA{132, 94, 194, 0xff}
	var i = 0
	for i < width {
		line, _, _ := rd.ReadLine()
		for l := 0; l < 10; l++ {
			j = 0
			for _, v := range line {
				switch v {
				case '0':
					for k := 0; k < 10; k++ {
						img.Set(j, i, color.White)
						j++
					}
				case '1':
					for k := 0; k < 10; k++ {
						img.Set(j, i, blue)
						j++
					}
				case '2':
					for k := 0; k < 10; k++ {
						img.Set(j, i, orange)
						j++
					}
				case '3':
					for k := 0; k < 10; k++ {
						img.Set(j, i, pink)
						j++
					}
				case '4':
					for k := 0; k < 10; k++ {
						img.Set(j, i, purple)
						j++
					}
				}
			}
			i++
		}
	}

	//var black = true
	//for i := 0; i < width; i++ {
	//	if i%10 == 0 {
	//		black = !black
	//	}
	//	for j := 0; j < height; j++ {
	//		if j%10 == 0 {
	//			black = !black
	//		}
	//		switch {
	//		case black:
	//			img.Set(j, i, color.Black)
	//		default:
	//			img.Set(j, i, color.White)
	//		}
	//	}
	//}
	return img
}

func CreateFile() {
	img := GenerateLogo()
	file, err := os.Create(FileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln(err)
	}
}
