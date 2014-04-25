package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel;
}

func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.Width, self.Height)
}

func (self Image)  At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), 128, 128, 255}
}

func main() {
	m := Image{ 200, 100 }
	pic.ShowImage(m)
}
