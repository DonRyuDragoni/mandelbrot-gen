package main

import (
	"fmt"

	"./image"
)

var (
	imgfile          = "img.png"
	size_y           = 2000
	size_x           = size_y * 16 / 9
	maxItersPerPixel = 500
)

func main() {
	img := image.NewMandelbrot(size_x, size_y, maxItersPerPixel,
		image.DefaultColorFn)

	err := img.Img.Encode(imgfile, true)
	if err != nil {
		fmt.Println("error:", err)
	}
}
