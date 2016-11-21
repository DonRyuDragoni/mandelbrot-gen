package image

import (
	"sync"
)

type MandelbrotSetImage struct {
	img    *Image
	width  int
	height int
}

/*
Default function used to generate the color of each pixel.

x                -> horizontal coordinate of the pixel
y                -> vertical coordinate of the pixel
itersDone        -> number of iterations
maxItersPerPixel -> maximum allowed number of iterations per pixel

colors -> array containing the RGBA values for the given pixel
*/
func DefaultColorFn(x, y, itersDone, maxItersPerPixel int) (colors [4]uint8) {
	if itersDone != maxItersPerPixel {
		brightness := uint8(remap(
			float32(itersDone),
			0, float32(maxItersPerPixel),
			0, 255,
		))
		colors[0] = brightness
		colors[1] = brightness
		colors[2] = brightness
		colors[3] = 255
	} else {
		colors[0] = 0
		colors[1] = 0
		colors[2] = 0
		colors[3] = 255
	}

	return
}

type ColorFn func(int, int, int, int) [4]uint8

func NewMandelbrot(size_x, size_y, maxItersPerPixel int,
	colorer ColorFn, numThreads int) (m *MandelbrotSetImage) {
	m = new(MandelbrotSetImage)
	m.width = size_x
	m.height = size_y
	m.img = NewImage(m.width, m.height)

	threadPoints := pointList(m.width, numThreads)

	var wg sync.WaitGroup

	for _, p := range threadPoints {
		wg.Add(1)

		go func(startx, endx int) {
			for x := startx; x <= endx; x++ {
				for y := 0; y <= m.height; y++ {
					n := pointTendsToinfinity(x, y, m.width, m.height,
						maxItersPerPixel)

					color := colorer(x, y, n, maxItersPerPixel)

					m.setPixel(x, y, color[0], color[1], color[2], color[3])
				}
			}

			wg.Done()
		}(p[0], p[1])
	}

	wg.Wait()

	return
}

func (m *MandelbrotSetImage) setPixel(x, y int, r, g, b, a uint8) {
	m.img.SetPixel(x, y, r, g, b, a)
}

func (m *MandelbrotSetImage) Encode(outfile string,
	overrideIfExists bool) error {

	return m.img.Encode(outfile, overrideIfExists)
}
