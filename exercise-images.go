package main

import (
	"golang.org/x/tour/pic"
	"math"
	"image"
	"image/color"
)

type Image struct {
	sx, sy, ex, ey int
}

func (i Image) Bounds() image.Rectangle {
	result := image.Rect(0, 0, int(math.Abs(float64(i.ex-i.sx))), int(math.Abs(float64(i.ey-i.sy))))
	return result
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.Gray{uint8(x^y)}
}

func main() {
	m := Image{0,0,100,100}
	pic.ShowImage(m)
}
