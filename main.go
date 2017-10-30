package main

import (
	"fmt"

	"github.com/DonRyuDragoni/mandelbrot/image"
)

var (
	imgfile          = "img.png"
	size_y           = 3000
	size_x           = size_y * 16 / 9
	maxItersPerPixel = 500
	threads          = 4
)

func ColorFn(x, y, itersDone, maxItersPerPixel int) (colors [4]uint8) {
	brightness := uint8(image.Remap(
		float32(itersDone),
		0, float32(maxItersPerPixel),
		70, 200,
	))

	colors[0] = brightness
	colors[1] = brightness
	colors[2] = brightness
	colors[3] = 255

	return
}

func main() {
	img := image.NewMandelbrot(size_x, size_y, maxItersPerPixel,
		ColorFn, threads)

	err := img.Encode(imgfile, true)
	if err != nil {
		fmt.Println("error:", err)
	}
}
