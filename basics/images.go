package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	w int
	h int
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x^y), uint8(x+y), 255,255}
}


func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.w,i.h)
}



func main() {
	m := Image{256, 64}
	pic.ShowImage(&m)
}
