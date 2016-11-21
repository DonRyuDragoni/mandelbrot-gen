package main

import (
	"fmt"

	"./image"
)

var (
	imgfile          = "img.png"
	size_y           = 3000
	size_x           = size_y * 16 / 9
	maxItersPerPixel = 500
	threads          = 4
)

func main() {
	img := image.NewMandelbrot(size_x, size_y, maxItersPerPixel,
		image.DefaultColorFn, threads)

	err := img.Encode(imgfile, true)
	if err != nil {
		fmt.Println("error:", err)
	}
}
