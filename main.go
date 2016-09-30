package main

import (
	"fmt"

	"./image"
)

var (
	imgfile          = "img.png"
	size_x           = 1000
	size_y           = 1000
	maxItersPerPixel = 100
)

func main() {
	img := image.NewMandelbrot(size_x, size_y, maxItersPerPixel,
		image.DefaultColorFn)

	err := img.Img.Encode(imgfile, true)
	if err != nil {
		fmt.Println("error:", err)
	}
}
